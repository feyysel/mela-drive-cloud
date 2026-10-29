package stringsx

import "strings"

// Trim removes leading and trailing whitespace from s.
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// TrimOptional trims *s. It returns nil if s is nil or blank.
func TrimOptional(s *string) *string {
	if s == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*s)
	if trimmed == "" {
		return nil
	}

	return &trimmed
}
