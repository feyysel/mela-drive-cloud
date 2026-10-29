package stringsx

import "strings"

// IsBlank reports whether s is empty or only whitespace.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsBlankPtr reports whether s is nil or blank.
func IsBlankPtr(s *string) bool {
	return s == nil || IsBlank(*s)
}

// PtrIfNotBlank returns &s, or nil if s is blank.
func PtrIfNotBlank(s string) *string {
	if IsBlank(s) {
		return nil
	}

	return &s
}
