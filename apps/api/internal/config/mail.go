package config

import "net/mail"

type MailConfig struct {
	APIKey string
	Domain string
}

func LoadMail() (MailConfig, error) {
	key, err := get("MAIL_API_KEY")
	if err != nil {
		return MailConfig{}, err
	}

	domain, err := get("MAIL_DOMAIN")
	if err != nil {
		return MailConfig{}, err
	}

	return MailConfig{
		APIKey: key,
		Domain: domain,
	}, nil
}

func (c MailConfig) NoReplyAddress(displayName string) string {
	return c.buildAddress(displayName, "no-reply")
}

func (c MailConfig) SecurityAddress(displayName string) string {
	return c.buildAddress(displayName, "security")
}

func (c MailConfig) buildAddress(displayName, localPart string) string {
	address := mail.Address{Name: displayName, Address: localPart + "@" + c.Domain}

	return address.String()
}
