package handlers

import (
	"net/http"

	"github.com/Nokeni/GODS/internal/web/api/services"
	"github.com/Nokeni/GODS/internal/web/common/dtos"
	"github.com/gin-gonic/gin"
)

// AuthHandler defines the interface for user-authentication-related HTTP handlers.
// @title UserHandler Interface
// @description Interface for handling user-authentication-related HTTP requests.
type AuthHandler interface {
	Login(c *gin.Context)
	Signup(c *gin.Context)
}

// UserHandlerImplementation handles HTTP requests for CRUD operations against the user model.
type AuthHandlerImplementation struct {
	authService services.AuthService
}

// NewUserHandler creates a new instance of the UserHandlerImplementation.
func NewAuthHandler(authService services.AuthService) *AuthHandlerImplementation {
	return &AuthHandlerImplementation{
		authService: authService,
	}
}

// Login authenticates a user.
// @Summary Authenticate a user
// @Description Authenticate a user with their username and password
// @Tags auth
// @Accept mpfd
// @Produce json
// @Param name formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} gin.H "JWT Token"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Router /auth/login [post]
func (handler *AuthHandlerImplementation) Login(c *gin.Context) {
	var loginDTO dtos.LoginDTO
	if err := c.ShouldBind(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := handler.authService.Login(&loginDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Signup creates a new user.
// @Summary Create a new user
// @Description Register a new user with username, password, and email
// @Tags auth
// @Accept mpfd
// @Produce json
// @Param name formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param password_confirmation formData string true "Password confirmation"
// @Success 201
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /auth/signup [post]
func (handler *AuthHandlerImplementation) Signup(c *gin.Context) {
	var signupDTO dtos.SignupDTO
	if err := c.ShouldBind(&signupDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.authService.Signup(&signupDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
