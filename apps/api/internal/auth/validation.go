package auth

import (
	"slices"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/feyysel/mela-drive-cloud/internal/emailx"
	"github.com/feyysel/mela-drive-cloud/internal/errs"
)

func requireValidEmail(email string) error {
	if err := emailx.Validate(email); err != nil {
		return errs.Wrap(errs.ErrInvalidInput, "Invalid email")
	}

	return nil
}

func requireValidOrigin(cfg config.AppConfig, origin string) error {
	if !slices.Contains(cfg.AllowedOrigins, origin) {
		return errs.Wrap(errs.ErrInternal, "invalid origin")
	}

	return nil
}
