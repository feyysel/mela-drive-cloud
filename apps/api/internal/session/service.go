package session

import (
	"context"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/google/uuid"
)

type Service struct {
	cache Cache
	cfg   config.SessionConfig
}

func NewService(cache Cache, cfg config.SessionConfig) *Service {
	return &Service{
		cache: cache,
		cfg:   cfg,
	}
}

func (s *Service) Create(ctx context.Context, userID uuid.UUID, ip, userAgent string) (*Session, error) {
	sess, err := NewSession(userID, ip, userAgent)
	if err != nil {
		return nil, err
	}

	if err := s.cache.Set(ctx, sessionKey(sess.ID), sess, s.cfg.TTL); err != nil {
		return nil, err
	}

	userSessKey := userSessionsKey(userID)

	if err := s.cache.SAdd(ctx, userSessKey, sess.ID.String()); err != nil {
		return nil, err
	}

	if err := s.cache.Expire(ctx, userSessKey, s.cfg.TTL); err != nil {
		return nil, err
	}

	return sess, nil
}

func (s *Service) Get(ctx context.Context, sessionID uuid.UUID) (*Session, error) {
	key := sessionKey(sessionID)

	sess, err := s.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	remaining, err := s.cache.TTL(ctx, key)
	if err == nil && remaining < s.cfg.RenewThreshold {
		_ = s.cache.Expire(ctx, key, s.cfg.TTL)
		_ = s.cache.Expire(ctx, userSessionsKey(sess.UserID), s.cfg.TTL)
	}

	return sess, nil
}

func (s *Service) Revoke(ctx context.Context, userID, sessionID uuid.UUID) error {
	if _, err := s.cache.Del(ctx, sessionKey(sessionID)); err != nil {
		return err
	}

	return s.cache.SRem(ctx, userSessionsKey(userID), sessionID.String())
}

func (s *Service) RevokeAll(ctx context.Context, userID uuid.UUID) error {
	userSessKey := userSessionsKey(userID)

	ids, err := s.cache.SMembers(ctx, userSessKey)
	if err != nil {
		return err
	}

	for _, id := range ids {
		sessionID, err := uuid.Parse(id)
		if err != nil {
			continue
		}

		if _, err := s.cache.Del(ctx, sessionKey(sessionID)); err != nil {
			return err
		}
	}

	if _, err := s.cache.Del(ctx, userSessKey); err != nil {
		return err
	}

	return nil
}
