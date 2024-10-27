package dtos

// LoginDTO represents the login credentials of a user.
type LoginDTO struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// SignupDTO represents the signup informations for a user.
type SignupDTO struct {
	Name                 string `form:"name" binding:"required"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" binding:"required"`
}
