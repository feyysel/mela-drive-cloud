package auth

import (
	"net/http"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/ginx"
	"github.com/feyysel/mela-drive-cloud/internal/res"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Middleware struct {
	sessionSvc SessionService
	userSvc    UserService

	sessionCfg config.SessionConfig
}

func NewMiddleware(sessionSvc SessionService, userSvc UserService, sessionCfg config.SessionConfig) *Middleware {
	return &Middleware{
		sessionSvc: sessionSvc,
		userSvc:    userSvc,
		sessionCfg: sessionCfg,
	}
}

func (m *Middleware) RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(m.sessionCfg.CookieName)
		if err != nil {
			m.clearSessionCookie(ctx)
			res.Error(ctx, errs.Wrap(errs.ErrUnauthorized, "Unauthorized"))

			return
		}

		sessionID, err := uuid.Parse(token)
		if err != nil {
			m.clearSessionCookie(ctx)
			res.Error(ctx, errs.Wrap(errs.ErrUnauthorized, "Unauthorized"))

			return
		}

		session, err := m.sessionSvc.Get(ctx, sessionID)
		if err != nil {
			m.clearSessionCookie(ctx)
			res.Error(ctx, errs.Wrap(errs.ErrUnauthorized, "Unauthorized"))

			return
		}

		userSnapshot, err := m.userSvc.GetSnapshot(ctx, session.UserID)
		if err != nil {
			m.clearSessionCookie(ctx)
			res.Error(ctx, errs.Wrap(errs.ErrUnauthorized, "Unauthorized"))

			return
		}

		ginx.SetSessionID(ctx, session.ID)
		ginx.SetUserContext(ctx, userSnapshot.ID, userSnapshot.Email, string(userSnapshot.Role))

		ctx.Next()
	}
}

func (m *Middleware) clearSessionCookie(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(m.sessionCfg.CookieName, "", -1, "/", m.sessionCfg.CookieDomain, m.sessionCfg.CookieSecure, true)
}
