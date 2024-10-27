package handlers

import (
	"net/http"
	"strconv"

	_ "github.com/Nokeni/GODS/internal/web/api/models"
	"github.com/Nokeni/GODS/internal/web/api/services"
	"github.com/Nokeni/GODS/internal/web/common/dtos"
	"github.com/gin-gonic/gin"
)

// UserHandler defines the interface for user-related HTTP handlers.
// @title UserHandler Interface
// @description Interface for handling user-related HTTP requests.
type UserHandler interface {
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// UserHandlerImplementation handles HTTP requests for CRUD operations against the user model.
type UserHandlerImplementation struct {
	userService services.UserService
}

// NewUserHandler creates a new instance of the UserHandlerImplementation.
func NewUserHandler(userService services.UserService) *UserHandlerImplementation {
	return &UserHandlerImplementation{
		userService: userService,
	}
}

// Get retrieves a user by ID.
// @Summary Get a user by ID
// @Description Get details of a user by their ID
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users/{id} [get]
func (handler *UserHandlerImplementation) Get(c *gin.Context) {
	id := c.Param("id")

	// Convert id from string to uint
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := handler.userService.Get(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAll retrieves all users.
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.User
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users [get]
func (handler *UserHandlerImplementation) GetAll(c *gin.Context) {
	users, err := handler.userService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Create adds a new user.
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept mpfd
// @Produce json
// @Security BearerAuth
// @Param name formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param password_confirmation formData string true "Password confirmation"
// @Success 201 {object} models.User
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users [post]
func (handler *UserHandlerImplementation) Create(c *gin.Context) {
	var userDTO dtos.CreateUserDTO
	if err := c.ShouldBind(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := handler.userService.Create(&userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Update modifies an existing user.
// @Summary Update an existing user
// @Description Update the details of an existing user by their ID
// @Tags users
// @Accept mpfd
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param name formData string false "Username"
// @Param email formData string false "Email"
// @Param password formData string false "Password"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users/{id} [put]
func (handler *UserHandlerImplementation) Update(c *gin.Context) {
	id := c.Param("id")

	// Convert id from string to uint
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var userDTO dtos.UpdateUserDTO
	if err := c.ShouldBind(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := handler.userService.Get(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := handler.userService.Update(user, &userDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete removes a user.
// @Summary Delete a user
// @Description Remove a user from the system
// @Tags users
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users/{id} [delete]
func (handler *UserHandlerImplementation) Delete(c *gin.Context) {
	id := c.Param("id")

	// Convert id from string to uint
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := handler.userService.Delete(uint(uid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
