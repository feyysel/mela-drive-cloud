package mail

import (
	"github.com/feyysel/mela-drive-cloud/internal/emailx"
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/stringsx"
)

func requireValidFrom(from string) error {
	if stringsx.IsBlank(from) {
		return errs.Wrap(errs.ErrInternal, "from is required")
	}

	if err := emailx.Validate(from); err != nil {
		return errs.Wrap(errs.ErrInternal, "invalid from address")
	}

	return nil
}

func requireValidRecipients(to []string) error {
	if len(to) == 0 {
		return errs.Wrap(errs.ErrInternal, "at least one recipient is required")
	}

	for _, recipient := range to {
		if stringsx.IsBlank(recipient) {
			return errs.Wrap(errs.ErrInternal, "recipient email is required")
		}

		if err := emailx.Validate(recipient); err != nil {
			return errs.Wrap(errs.ErrInternal, "invalid recipient email")
		}
	}

	return nil
}

func requireValidTemplateID(templateID string) error {
	if stringsx.IsBlank(templateID) {
		return errs.Wrap(errs.ErrInternal, "template id is required")
	}

	return nil
}
