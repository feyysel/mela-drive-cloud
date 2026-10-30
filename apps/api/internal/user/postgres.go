package user

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

const (
	columns = "id, first_name, last_name, email, password_hash, role, created_at, updated_at"

	constraintEmail = "users_email_key"
)

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) Create(ctx context.Context, user *User) (*User, error) {
	query := fmt.Sprintf("INSERT INTO users (id, first_name, last_name, email, password_hash, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING %s", columns)

	created, err := scanUser(r.pool.QueryRow(ctx, query, user.ID, user.FirstName, user.LastName, user.Email, user.PasswordHash, user.Role))
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == pgerrcode.UniqueViolation && pgErr.ConstraintName == constraintEmail {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return created, nil
}

func (r *postgresRepository) GetByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM users WHERE id = $1", columns)

	user, err := scanUser(r.pool.QueryRow(ctx, query, userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (r *postgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM users WHERE email = $1", columns)

	user, err := scanUser(r.pool.QueryRow(ctx, query, email))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (r *postgresRepository) CountByRole(ctx context.Context, role UserRole) (int, error) {
	query := "SELECT COUNT(*) FROM users WHERE role = $1"

	var count int
	if err := r.pool.QueryRow(ctx, query, role).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *postgresRepository) UpdateName(ctx context.Context, userID uuid.UUID, firstName string, lastName *string) (*User, error) {
	query := fmt.Sprintf("UPDATE users SET first_name = $1, last_name = $2, updated_at = now() WHERE id = $3 RETURNING %s", columns)

	updated, err := scanUser(r.pool.QueryRow(ctx, query, firstName, lastName, userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return updated, nil
}

func (r *postgresRepository) UpdateEmail(ctx context.Context, userID uuid.UUID, email string) (*User, error) {
	query := fmt.Sprintf("UPDATE users SET email = $1, updated_at = now() WHERE id = $2 RETURNING %s", columns)

	updated, err := scanUser(r.pool.QueryRow(ctx, query, email, userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == pgerrcode.UniqueViolation && pgErr.ConstraintName == constraintEmail {
			return nil, errs.ErrConflict
		}

		return nil, err
	}

	return updated, nil
}

func (r *postgresRepository) UpdatePassword(ctx context.Context, userID uuid.UUID, passwordHash string) (*User, error) {
	query := fmt.Sprintf("UPDATE users SET password_hash = $1, updated_at = now() WHERE id = $2 RETURNING %s", columns)

	updated, err := scanUser(r.pool.QueryRow(ctx, query, passwordHash, userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	return updated, nil
}

func (r *postgresRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1"

	tag, err := r.pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func scanUser(row pgx.Row) (*User, error) {
	var user User

	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
