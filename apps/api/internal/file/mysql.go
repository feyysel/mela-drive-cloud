package file

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

const columns = "id, owner_id, folder_id, name, storage_key, size_bytes, mime_type, deleted_at, created_at, updated_at"

type scanner interface {
	Scan(dest ...any) error
}

type mysqlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) Create(ctx context.Context, file *File) (*File, error) {
	query := "INSERT INTO files (id, owner_id, folder_id, name, storage_key, size_bytes, mime_type) VALUES (?, ?, ?, ?, ?, ?, ?)"

	if _, err := r.db.ExecContext(ctx, query, file.ID, file.OwnerID, file.FolderID, file.Name, file.StorageKey, file.SizeBytes, file.MimeType); err != nil {
		if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok && mysqlErr.Number == mysqlerr.ER_DUP_ENTRY {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return r.Get(ctx, file.OwnerID, file.ID)
}

func (r *mysqlRepository) Get(ctx context.Context, ownerID, fileID uuid.UUID) (*File, error) {
	query := fmt.Sprintf("SELECT %s FROM files WHERE id = ? AND owner_id = ? AND deleted_at IS NULL", columns)

	file, err := scanFile(r.db.QueryRowContext(ctx, query, fileID, ownerID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return file, nil
}

func (r *mysqlRepository) List(ctx context.Context, ownerID uuid.UUID, folderID *uuid.UUID) ([]*File, error) {
	query := fmt.Sprintf("SELECT %s FROM files WHERE owner_id = ? AND folder_id <=> ? AND deleted_at IS NULL ORDER BY name", columns)

	rows, err := r.db.QueryContext(ctx, query, ownerID, folderID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	files := []*File{}

	for rows.Next() {
		file, err := scanFile(rows)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *mysqlRepository) Rename(ctx context.Context, ownerID, fileID uuid.UUID, name string) (*File, error) {
	query := "UPDATE files SET name = ? WHERE id = ? AND owner_id = ? AND deleted_at IS NULL"

	if _, err := r.db.ExecContext(ctx, query, name, fileID, ownerID); err != nil {
		if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok && mysqlErr.Number == mysqlerr.ER_DUP_ENTRY {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return r.Get(ctx, ownerID, fileID)
}

func (r *mysqlRepository) Move(ctx context.Context, ownerID, fileID uuid.UUID, newFolderID *uuid.UUID) (*File, error) {
	query := "UPDATE files SET folder_id = ? WHERE id = ? AND owner_id = ? AND deleted_at IS NULL"

	if _, err := r.db.ExecContext(ctx, query, newFolderID, fileID, ownerID); err != nil {
		if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok && mysqlErr.Number == mysqlerr.ER_DUP_ENTRY {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return r.Get(ctx, ownerID, fileID)
}

func (r *mysqlRepository) GetTrashed(ctx context.Context, ownerID, fileID uuid.UUID) (*File, error) {
	query := fmt.Sprintf("SELECT %s FROM files WHERE id = ? AND owner_id = ? AND deleted_at IS NOT NULL", columns)

	file, err := scanFile(r.db.QueryRowContext(ctx, query, fileID, ownerID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return file, nil
}

func (r *mysqlRepository) ListTrashed(ctx context.Context, ownerID uuid.UUID) ([]*File, error) {
	query := fmt.Sprintf(`SELECT %s FROM files
		WHERE owner_id = ? AND deleted_at IS NOT NULL
		AND NOT EXISTS (SELECT 1 FROM folders WHERE folders.id = files.folder_id AND folders.deleted_at IS NOT NULL)
		ORDER BY deleted_at DESC`, columns)

	rows, err := r.db.QueryContext(ctx, query, ownerID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	files := []*File{}

	for rows.Next() {
		file, err := scanFile(rows)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *mysqlRepository) Trash(ctx context.Context, ownerID, fileID uuid.UUID) error {
	query := "UPDATE files SET deleted_at = CURRENT_TIMESTAMP(6) WHERE id = ? AND owner_id = ? AND deleted_at IS NULL"

	result, err := r.db.ExecContext(ctx, query, fileID, ownerID)
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

func (r *mysqlRepository) Restore(ctx context.Context, ownerID, fileID uuid.UUID, folderID *uuid.UUID) error {
	query := "UPDATE files SET deleted_at = NULL, folder_id = ? WHERE id = ? AND owner_id = ? AND deleted_at IS NOT NULL"

	result, err := r.db.ExecContext(ctx, query, folderID, fileID, ownerID)
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

func (r *mysqlRepository) Delete(ctx context.Context, ownerID, fileID uuid.UUID) error {
	query := "DELETE FROM files WHERE id = ? AND owner_id = ? AND deleted_at IS NOT NULL"

	result, err := r.db.ExecContext(ctx, query, fileID, ownerID)
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

func scanFile(row scanner) (*File, error) {
	var file File

	if err := row.Scan(&file.ID, &file.OwnerID, &file.FolderID, &file.Name, &file.StorageKey, &file.SizeBytes, &file.MimeType, &file.DeletedAt, &file.CreatedAt, &file.UpdatedAt); err != nil {
		return nil, err
	}

	return &file, nil
}
