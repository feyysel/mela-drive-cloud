package user

import (
	"unicode/utf8"

	"github.com/feyysel/mela-drive-cloud/internal/emailx"
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/stringsx"
)

func requireValidFirstName(firstName string) error {
	if firstName == "" {
		return errs.Wrap(errs.ErrInvalidInput, "First name is required")
	}

	length := utf8.RuneCountInString(firstName)

	if length < 2 {
		return errs.Wrap(errs.ErrInvalidInput, "First name is too short")
	}

	if length > 50 {
		return errs.Wrap(errs.ErrInvalidInput, "First name is too long")
	}

	return nil
}

func requireValidLastName(lastName *string) error {
	if lastName == nil {
		return nil
	}

	if utf8.RuneCountInString(*lastName) > 50 {
		return errs.Wrap(errs.ErrInvalidInput, "Last name is too long")
	}

	return nil
}

func requireValidEmail(email string) error {
	if email == "" {
		return errs.Wrap(errs.ErrInvalidInput, "Email is required")
	}

	if err := emailx.Validate(email); err != nil {
		return errs.Wrap(errs.ErrInvalidInput, "Invalid email")
	}

	return nil
}

func requireValidPassword(password string) error {
	if stringsx.IsBlank(password) {
		return errs.Wrap(errs.ErrInvalidInput, "Password is required")
	}

	if utf8.RuneCountInString(password) < 8 {
		return errs.Wrap(errs.ErrInvalidInput, "Password must be at least 8 characters")
	}

	return nil
}

func requireValidPasswordHash(passwordHash string) error {
	if passwordHash == "" {
		return errs.Wrap(errs.ErrInvalidInput, "Password is required")
	}

	return nil
}

func requireValidRole(role UserRole) error {
	if !role.IsValid() {
		return errs.Wrap(errs.ErrInvalidInput, "Invalid role")
	}

	return nil
}
