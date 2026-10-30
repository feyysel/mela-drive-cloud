package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/session"
	"github.com/feyysel/mela-drive-cloud/internal/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const msgInvalidCredentials = "Invalid email or password"

type Service struct {
	cache      Cache
	userSvc    UserService
	sessionSvc SessionService
	mailSvc    MailService
	appCfg     config.AppConfig
	mailCfg    config.MailConfig
}

func NewService(cache Cache, userSvc UserService, sessionSvc SessionService, mailSvc MailService, appCfg config.AppConfig, mailCfg config.MailConfig) *Service {
	return &Service{
		cache:      cache,
		userSvc:    userSvc,
		sessionSvc: sessionSvc,
		mailSvc:    mailSvc,
		appCfg:     appCfg,
		mailCfg:    mailCfg,
	}
}

func (s *Service) Signup(ctx context.Context, firstName string, lastName *string, email, password, ip, userAgent string) (*session.Session, error) {
	u, err := s.userSvc.Create(ctx, firstName, lastName, email, password, user.UserRoleMember)
	if err != nil {
		return nil, err
	}

	return s.sessionSvc.Create(ctx, u.ID, ip, userAgent)
}

func (s *Service) Login(ctx context.Context, email, password, ip, userAgent string) (*session.Session, error) {
	u, err := s.userSvc.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrUnauthorized, msgInvalidCredentials)
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return nil, errs.Wrap(errs.ErrUnauthorized, msgInvalidCredentials)
	}

	return s.sessionSvc.Create(ctx, u.ID, ip, userAgent)
}

func (s *Service) Logout(ctx context.Context, userID, sessionID uuid.UUID) error {
	return s.sessionSvc.Revoke(ctx, userID, sessionID)
}

func (s *Service) ForgotPassword(ctx context.Context, email, origin string) error {
	if err := requireValidOrigin(s.appCfg, origin); err != nil {
		return err
	}

	u, err := s.userSvc.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil
		}

		return err
	}

	token := rand.Text()
	sum := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(sum[:])

	if err := s.cache.Set(ctx, resetTokenKey(tokenHash), u.ID, resetTokenTTL); err != nil {
		return err
	}

	return s.mailSvc.SendTemplateEmail(
		s.mailCfg.SecurityAddress(s.appCfg.Name),
		[]string{u.Email},
		"reset-password-email-template",
		map[string]any{
			"FIRST_NAME": u.FirstName,
			"RESET_LINK": origin + "/reset-password?token=" + token,
		},
	)
}

func (s *Service) ResetPassword(ctx context.Context, token, password string) error {
	sum := sha256.Sum256([]byte(token))
	key := resetTokenKey(hex.EncodeToString(sum[:]))

	userID, err := s.cache.Get(ctx, key)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.Wrap(errs.ErrInvalidInput, "Invalid or expired reset token")
		}

		return err
	}

	if _, err := s.userSvc.UpdatePassword(ctx, userID, password); err != nil {
		return err
	}

	if err := s.cache.Del(ctx, key); err != nil {
		return err
	}

	return s.sessionSvc.RevokeAll(ctx, userID)
}
