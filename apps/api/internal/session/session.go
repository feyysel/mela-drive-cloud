package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"userAgent"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewSession(userID uuid.UUID, ip, userAgent string) (*Session, error) {
	if err := requireValidUserID(userID); err != nil {
		return nil, err
	}

	if err := requireValidIP(ip); err != nil {
		return nil, err
	}

	if err := requireValidUserAgent(userAgent); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:        id,
		UserID:    userID,
		IP:        ip,
		UserAgent: userAgent,
		CreatedAt: time.Now(),
	}, nil
}
