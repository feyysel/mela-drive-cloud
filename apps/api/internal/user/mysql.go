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
	columns = "id, first_name, last_name, email, password_hash, role, created_at, updated_at"

	keyEmail = "users.email"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) Create(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO users (id, first_name, last_name, email, password_hash, role) VALUES (?, ?, ?, ?, ?, ?)"

	if _, err := r.db.ExecContext(ctx, query, user.ID, user.FirstName, user.LastName, user.Email, user.PasswordHash, user.Role); err != nil {
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

func (r *mysqlRepository) CountByRole(ctx context.Context, role UserRole) (int, error) {
	query := "SELECT COUNT(*) FROM users WHERE role = ?"

	var count int
	if err := r.db.QueryRowContext(ctx, query, role).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *mysqlRepository) UpdateName(ctx context.Context, userID uuid.UUID, firstName string, lastName *string) (*User, error) {
	query := "UPDATE users SET first_name = ?, last_name = ? WHERE id = ?"

	if _, err := r.db.ExecContext(ctx, query, firstName, lastName, userID); err != nil {
		return nil, err
	}

	return r.GetByID(ctx, userID)
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

	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
