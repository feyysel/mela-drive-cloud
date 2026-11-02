package auth

import (
	"net/http"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/feyysel/mela-drive-cloud/internal/ginx"
	"github.com/feyysel/mela-drive-cloud/internal/res"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc        *Service
	sessionCfg config.SessionConfig
}

func NewHandler(svc *Service, sessionCfg config.SessionConfig) *Handler {
	return &Handler{
		svc:        svc,
		sessionCfg: sessionCfg,
	}
}

func (h *Handler) Signup(ctx *gin.Context) {
	var req SignupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	sess, err := h.svc.Signup(ctx, req.FirstName, req.LastName, req.Email, req.Password, ctx.ClientIP(), ctx.Request.UserAgent())
	if err != nil {
		res.Error(ctx, err)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(h.sessionCfg.CookieName, sess.ID.String(), h.sessionCfg.CookieMaxAge, "/", h.sessionCfg.CookieDomain, h.sessionCfg.CookieSecure, true)

	res.NoContent(ctx)
}

func (h *Handler) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	sess, err := h.svc.Login(ctx, req.Email, req.Password, ctx.ClientIP(), ctx.Request.UserAgent())
	if err != nil {
		res.Error(ctx, err)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(h.sessionCfg.CookieName, sess.ID.String(), h.sessionCfg.CookieMaxAge, "/", h.sessionCfg.CookieDomain, h.sessionCfg.CookieSecure, true)

	res.NoContent(ctx)
}

func (h *Handler) Logout(ctx *gin.Context) {
	userID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	sessionID, err := ginx.SessionID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	if err := h.svc.Logout(ctx, userID, sessionID); err != nil {
		res.Error(ctx, err)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(h.sessionCfg.CookieName, "", -1, "/", h.sessionCfg.CookieDomain, h.sessionCfg.CookieSecure, true)

	res.NoContent(ctx)
}

func (h *Handler) ForgotPassword(ctx *gin.Context) {
	var req ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	if err := h.svc.ForgotPassword(ctx, req.Email, ctx.GetHeader("Origin")); err != nil {
		res.Error(ctx, err)
		return
	}

	res.NoContent(ctx)
}

func (h *Handler) ResetPassword(ctx *gin.Context) {
	var req ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	if err := h.svc.ResetPassword(ctx, req.Token, req.Password); err != nil {
		res.Error(ctx, err)
		return
	}

	res.NoContent(ctx)
}

func (h *Handler) ValidateResetToken(ctx *gin.Context) {
	var req ValidateResetTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	if err := h.svc.ValidateResetToken(ctx, req.Token); err != nil {
		res.Error(ctx, err)
		return
	}

	res.NoContent(ctx)
}
