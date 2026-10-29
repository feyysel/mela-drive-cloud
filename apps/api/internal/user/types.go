package user

type CreateUserRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  *string `json:"lastName"`
	Email     string  `json:"email" binding:"required"`
}

type UpdateUserNameRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  *string `json:"lastName"`
}

type UpdateEmailRequest struct {
	Email string `json:"email" binding:"required"`
}
