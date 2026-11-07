package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/VividCortex/mysqlerr"
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

const (
	columns = "id, email, password_hash, storage_used_bytes, storage_quota_bytes, created_at, updated_at"

	keyEmail = "users.email"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) Create(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO users (id, email, password_hash) VALUES (?, ?, ?)"

	if _, err := r.db.ExecContext(ctx, query, user.ID, user.Email, user.PasswordHash); err != nil {
		if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok && mysqlErr.Number == mysqlerr.ER_DUP_ENTRY && strings.Contains(mysqlErr.Message, keyEmail) {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return r.GetByID(ctx, user.ID)
}

func (r *mysqlRepository) GetByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM users WHERE id = ?", columns)

	user, err := scanUser(r.db.QueryRowContext(ctx, query, userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (r *mysqlRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM users WHERE email = ?", columns)

	user, err := scanUser(r.db.QueryRowContext(ctx, query, email))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (r *mysqlRepository) UpdateEmail(ctx context.Context, userID uuid.UUID, email string) (*User, error) {
	query := "UPDATE users SET email = ? WHERE id = ?"

	if _, err := r.db.ExecContext(ctx, query, email, userID); err != nil {
		if mysqlErr, ok := errors.AsType[*mysql.MySQLError](err); ok && mysqlErr.Number == mysqlerr.ER_DUP_ENTRY && strings.Contains(mysqlErr.Message, keyEmail) {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return r.GetByID(ctx, userID)
}

func (r *mysqlRepository) UpdatePassword(ctx context.Context, userID uuid.UUID, passwordHash string) (*User, error) {
	query := "UPDATE users SET password_hash = ? WHERE id = ?"

	if _, err := r.db.ExecContext(ctx, query, passwordHash, userID); err != nil {
		return nil, err
	}

	return r.GetByID(ctx, userID)
}

func (r *mysqlRepository) ReserveStorage(ctx context.Context, userID uuid.UUID, sizeBytes int64) error {
	query := "UPDATE users SET storage_used_bytes = storage_used_bytes + ? WHERE id = ? AND storage_used_bytes + ? <= storage_quota_bytes"

	result, err := r.db.ExecContext(ctx, query, sizeBytes, userID, sizeBytes)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		if _, err := r.GetByID(ctx, userID); err != nil {
			return err
		}

		return errs.ErrConflict
	}

	return nil
}

func (r *mysqlRepository) ReleaseStorage(ctx context.Context, userID uuid.UUID, sizeBytes int64) error {
	query := "UPDATE users SET storage_used_bytes = GREATEST(storage_used_bytes - ?, 0) WHERE id = ?"

	_, err := r.db.ExecContext(ctx, query, sizeBytes, userID)

	return err
}

func (r *mysqlRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := r.db.ExecContext(ctx, query, userID)
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

func scanUser(row *sql.Row) (*User, error) {
	var user User

	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.StorageUsedBytes, &user.StorageQuotaBytes, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
