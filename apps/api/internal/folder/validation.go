package folder

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
)

const maxNameLength = 255

func requireValidName(name string) error {
	if name == "" {
		return errs.Wrap(errs.ErrInvalidInput, "Folder name is required")
	}

	if utf8.RuneCountInString(name) > maxNameLength {
		return errs.Wrap(errs.ErrInvalidInput, "Folder name is too long")
	}

	if name == "." || name == ".." || strings.ContainsFunc(name, isForbiddenNameRune) {
		return errs.Wrap(errs.ErrInvalidInput, "Folder name contains invalid characters")
	}

	return nil
}

func isForbiddenNameRune(r rune) bool {
	return r == '/' || r == '\\' || unicode.IsControl(r)
}
