package dtos

// CreateUserDTO represents the signup informations for a user.
type CreateUserDTO struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

// UpdateUserDTO represents the update informations for a user.
type UpdateUserDTO struct {
	Name     string `form:"name"`
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
}
