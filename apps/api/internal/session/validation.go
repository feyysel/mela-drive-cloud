package session

import (
	"net"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/stringsx"
	"github.com/google/uuid"
)

func requireValidUserID(userID uuid.UUID) error {
	if userID == uuid.Nil {
		return errs.Wrap(errs.ErrInternal, "user id is required")
	}

	return nil
}

func requireValidIP(ip string) error {
	if stringsx.IsBlank(ip) {
		return errs.Wrap(errs.ErrInternal, "ip is required")
	}

	if net.ParseIP(ip) == nil {
		return errs.Wrap(errs.ErrInternal, "invalid ip")
	}

	return nil
}

func requireValidUserAgent(userAgent string) error {
	if stringsx.IsBlank(userAgent) {
		return errs.Wrap(errs.ErrInternal, "user agent is required")
	}

	return nil
}
