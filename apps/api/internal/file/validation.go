package file

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
)

const (
	maxNameLength       = 255
	maxFileSize         = 100 << 20
	maxStorageKeyLength = 512
)

func requireValidName(name string) error {
	if name == "" {
		return errs.Wrap(errs.ErrInvalidInput, "File name is required")
	}

	if utf8.RuneCountInString(name) > maxNameLength {
		return errs.Wrap(errs.ErrInvalidInput, "File name is too long")
	}

	if name == "." || name == ".." || strings.ContainsFunc(name, isForbiddenNameRune) {
		return errs.Wrap(errs.ErrInvalidInput, "File name contains invalid characters")
	}

	return nil
}

func requireValidStorageKey(storageKey string) error {
	if storageKey == "" {
		return errs.Wrap(errs.ErrInvalidInput, "Storage key is required")
	}

	if len(storageKey) > maxStorageKeyLength {
		return errs.Wrap(errs.ErrInvalidInput, "Storage key is too long")
	}

	return nil
}

func requireValidSize(sizeBytes int64) error {
	if sizeBytes < 0 {
		return errs.Wrap(errs.ErrInvalidInput, "Invalid file size")
	}

	if sizeBytes > maxFileSize {
		return errs.Wrap(errs.ErrInvalidInput, "File is too large")
	}

	return nil
}

func requireValidMimeType(mimeType string) error {
	if mimeType == "" {
		return errs.Wrap(errs.ErrInvalidInput, "File type is required")
	}

	return nil
}

func isForbiddenNameRune(char rune) bool {
	return char == '/' || char == '\\' || unicode.IsControl(char)
}
