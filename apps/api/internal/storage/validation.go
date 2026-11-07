package storage

import (
	"strings"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/stringsx"
)

func requireValidKey(key string) error {
	if stringsx.IsBlank(key) {
		return errs.Wrap(errs.ErrInternal, "key is required")
	}

	return nil
}

func requireValidPrefix(prefix string) error {
	if stringsx.IsBlank(prefix) {
		return errs.Wrap(errs.ErrInternal, "prefix is required")
	}

	if !strings.HasSuffix(prefix, "/") {
		return errs.Wrap(errs.ErrInternal, "prefix must end with /")
	}

	return nil
}

func requireValidContentType(contentType string) error {
	if stringsx.IsBlank(contentType) {
		return errs.Wrap(errs.ErrInternal, "content type is required")
	}

	return nil
}
