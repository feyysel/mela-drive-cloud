package folder

import (
	"context"
	"errors"
	"fmt"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const columns = "id, owner_id, parent_id, name, deleted_at, created_at, updated_at"

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) Create(ctx context.Context, folder *Folder) (*Folder, error) {
	query := fmt.Sprintf("INSERT INTO folders (id, owner_id, parent_id, name) VALUES ($1, $2, $3, $4) RETURNING %s", columns)

	created, err := scanFolder(r.pool.QueryRow(ctx, query, folder.ID, folder.OwnerID, folder.ParentID, folder.Name))
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return created, nil
}

func (r *postgresRepository) Get(ctx context.Context, ownerID, folderID uuid.UUID) (*Folder, error) {
	query := fmt.Sprintf("SELECT %s FROM folders WHERE id = $1 AND owner_id = $2 AND deleted_at IS NULL", columns)

	folder, err := scanFolder(r.pool.QueryRow(ctx, query, folderID, ownerID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return folder, nil
}

func (r *postgresRepository) List(ctx context.Context, ownerID uuid.UUID, parentID *uuid.UUID) ([]*Folder, error) {
	query := fmt.Sprintf("SELECT %s FROM folders WHERE owner_id = $1 AND parent_id IS NOT DISTINCT FROM $2 AND deleted_at IS NULL ORDER BY lower(name)", columns)

	rows, err := r.pool.Query(ctx, query, ownerID, parentID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (*Folder, error) {
		return scanFolder(row)
	})
}

func (r *postgresRepository) Rename(ctx context.Context, ownerID, folderID uuid.UUID, name string) (*Folder, error) {
	query := fmt.Sprintf("UPDATE folders SET name = $1, updated_at = now() WHERE id = $2 AND owner_id = $3 AND deleted_at IS NULL RETURNING %s", columns)

	renamed, err := scanFolder(r.pool.QueryRow(ctx, query, name, folderID, ownerID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return renamed, nil
}

func (r *postgresRepository) Move(ctx context.Context, ownerID, folderID uuid.UUID, newParentID *uuid.UUID) (*Folder, error) {
	query := fmt.Sprintf("UPDATE folders SET parent_id = $1, updated_at = now() WHERE id = $2 AND owner_id = $3 AND deleted_at IS NULL RETURNING %s", columns)

	moved, err := scanFolder(r.pool.QueryRow(ctx, query, newParentID, folderID, ownerID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return moved, nil
}

func scanFolder(row pgx.Row) (*Folder, error) {
	var folder Folder

	if err := row.Scan(&folder.ID, &folder.OwnerID, &folder.ParentID, &folder.Name, &folder.DeletedAt, &folder.CreatedAt, &folder.UpdatedAt); err != nil {
		return nil, err
	}

	return &folder, nil
}
