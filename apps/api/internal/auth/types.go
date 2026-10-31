package auth

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  *string `json:"lastName"`
	Email     string  `json:"email" binding:"required"`
	Password  string  `json:"password" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ValidateResetTokenRequest struct {
	Token string `json:"token" binding:"required"`
}
