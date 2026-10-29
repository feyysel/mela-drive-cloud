package emailx

import (
	"net/mail"
)

// Validate returns an error if email is not a valid address.
func Validate(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
