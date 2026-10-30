package mail

import (
	"github.com/resend/resend-go/v4"
)

type Service struct {
	client *resend.Client
}

func NewService(client *resend.Client) *Service {
	return &Service{client: client}
}

func (s *Service) SendTemplateEmail(from string, to []string, templateID string, variables map[string]any) error {
	if err := requireValidFrom(from); err != nil {
		return err
	}

	if err := requireValidRecipients(to); err != nil {
		return err
	}

	if err := requireValidTemplateID(templateID); err != nil {
		return err
	}

	_, err := s.client.Emails.Send(&resend.SendEmailRequest{
		From: from,
		To:   to,
		Template: &resend.EmailTemplate{
			Id:        templateID,
			Variables: variables,
		},
	})

	return err
}
