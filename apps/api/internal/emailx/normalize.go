package emailx

import "strings"

// Normalize returns email in lowercase.
func Normalize(email string) string {
	return strings.ToLower(email)
}
