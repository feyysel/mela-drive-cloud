package emailx

import "strings"

// Mask returns email as a***@example.com, or *** if it can't be masked.
func Mask(email string) string {
	at := strings.IndexByte(email, '@')
	if at <= 1 {
		return "***"
	}

	return email[:1] + "***" + email[at:]
}
