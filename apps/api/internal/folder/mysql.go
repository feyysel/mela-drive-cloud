package folder

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/VividCortex/mysqlerr"
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const columns = "id, owner_id, parent_id, name, deleted_at, created_at, updated_at"

type scanner interface {
	Scan(dest ...any) error
}

type mysqlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) Create(ctx context.Context, folder *Folder) (*Folder, error) {
	query := "INSERT INTO folders (id, owner_id, parent_id, name) VALUES (?, ?, ?, ?)"

	if _, err := r.db.ExecContext(ctx, query, folder.ID, folder.OwnerID, folder.ParentID, folder.Name); err != nil {
		if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok && mysqlErr.Number == mysqlerr.ER_DUP_ENTRY {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return r.Get(ctx, folder.OwnerID, folder.ID)
}

func (r *mysqlRepository) Get(ctx context.Context, ownerID, folderID uuid.UUID) (*Folder, error) {
	query := fmt.Sprintf("SELECT %s FROM folders WHERE id = ? AND owner_id = ? AND deleted_at IS NULL", columns)

	folder, err := scanFolder(r.db.QueryRowContext(ctx, query, folderID, ownerID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return folder, nil
}

func (r *mysqlRepository) List(ctx context.Context, ownerID uuid.UUID, parentID *uuid.UUID) ([]*Folder, error) {
	query := fmt.Sprintf("SELECT %s FROM folders WHERE owner_id = ? AND parent_id <=> ? AND deleted_at IS NULL ORDER BY name", columns)

	rows, err := r.db.QueryContext(ctx, query, ownerID, parentID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	folders := []*Folder{}

	for rows.Next() {
		folder, err := scanFolder(rows)
		if err != nil {
			return nil, err
		}

		folders = append(folders, folder)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return folders, nil
}

func (r *mysqlRepository) Rename(ctx context.Context, ownerID, folderID uuid.UUID, name string) (*Folder, error) {
	query := "UPDATE folders SET name = ? WHERE id = ? AND owner_id = ? AND deleted_at IS NULL"

	if _, err := r.db.ExecContext(ctx, query, name, folderID, ownerID); err != nil {
		if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok && mysqlErr.Number == mysqlerr.ER_DUP_ENTRY {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return r.Get(ctx, ownerID, folderID)
}

func (r *mysqlRepository) Move(ctx context.Context, ownerID, folderID uuid.UUID, newParentID *uuid.UUID) (*Folder, error) {
	query := "UPDATE folders SET parent_id = ? WHERE id = ? AND owner_id = ? AND deleted_at IS NULL"

	if _, err := r.db.ExecContext(ctx, query, newParentID, folderID, ownerID); err != nil {
		if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok && mysqlErr.Number == mysqlerr.ER_DUP_ENTRY {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return r.Get(ctx, ownerID, folderID)
}

func (r *mysqlRepository) GetTrashed(ctx context.Context, ownerID, folderID uuid.UUID) (*Folder, error) {
	query := fmt.Sprintf("SELECT %s FROM folders WHERE id = ? AND owner_id = ? AND deleted_at IS NOT NULL", columns)

	folder, err := scanFolder(r.db.QueryRowContext(ctx, query, folderID, ownerID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return folder, nil
}

func (r *mysqlRepository) ListTrashed(ctx context.Context, ownerID uuid.UUID) ([]*Folder, error) {
	query := fmt.Sprintf(`SELECT %s FROM folders
		WHERE owner_id = ? AND deleted_at IS NOT NULL
		AND NOT EXISTS (SELECT 1 FROM folders AS parent WHERE parent.id = folders.parent_id AND parent.deleted_at IS NOT NULL)
		ORDER BY deleted_at DESC`, columns)

	rows, err := r.db.QueryContext(ctx, query, ownerID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	folders := []*Folder{}

	for rows.Next() {
		folder, err := scanFolder(rows)
		if err != nil {
			return nil, err
		}

		folders = append(folders, folder)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return folders, nil
}

func (r *mysqlRepository) Trash(ctx context.Context, ownerID, folderID uuid.UUID) error {
	query := "UPDATE folders SET deleted_at = CURRENT_TIMESTAMP(6) WHERE id = ? AND owner_id = ? AND deleted_at IS NULL"

	result, err := r.db.ExecContext(ctx, query, folderID, ownerID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *mysqlRepository) Restore(ctx context.Context, ownerID, folderID uuid.UUID, parentID *uuid.UUID) error {
	query := "UPDATE folders SET deleted_at = NULL, parent_id = ? WHERE id = ? AND owner_id = ? AND deleted_at IS NOT NULL"

	result, err := r.db.ExecContext(ctx, query, parentID, folderID, ownerID)
	if err != nil {
		if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok && mysqlErr.Number == mysqlerr.ER_DUP_ENTRY {
			return errs.ErrConflict
		}

		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *mysqlRepository) Delete(ctx context.Context, ownerID, folderID uuid.UUID) error {
	query := "DELETE FROM folders WHERE id = ? AND owner_id = ? AND deleted_at IS NOT NULL"

	result, err := r.db.ExecContext(ctx, query, folderID, ownerID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *mysqlRepository) ListFileKeys(ctx context.Context, ownerID, folderID uuid.UUID) ([]string, int64, error) {
	query := `WITH RECURSIVE tree AS (
		SELECT id FROM folders WHERE id = ? AND owner_id = ?
		UNION ALL
		SELECT f.id FROM folders f JOIN tree ON f.parent_id = tree.id WHERE f.owner_id = ?
	) SELECT storage_key, size_bytes FROM files WHERE owner_id = ? AND folder_id IN (SELECT id FROM tree)`

	rows, err := r.db.QueryContext(ctx, query, folderID, ownerID, ownerID, ownerID)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		_ = rows.Close()
	}()

	keys := []string{}

	var totalSize int64

	for rows.Next() {
		var key string
		var size int64

		if err := rows.Scan(&key, &size); err != nil {
			return nil, 0, err
		}

		keys = append(keys, key)
		totalSize += size
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return keys, totalSize, nil
}

func scanFolder(row scanner) (*Folder, error) {
	var folder Folder

	if err := row.Scan(&folder.ID, &folder.OwnerID, &folder.ParentID, &folder.Name, &folder.DeletedAt, &folder.CreatedAt, &folder.UpdatedAt); err != nil {
		return nil, err
	}

	return &folder, nil
}
