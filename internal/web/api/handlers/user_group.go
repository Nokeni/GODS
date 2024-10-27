package handlers

import (
	"net/http"
	"strconv"

	_ "github.com/Nokeni/GODS/internal/web/api/models"
	"github.com/Nokeni/GODS/internal/web/api/services"
	"github.com/gin-gonic/gin"
)

// UserGroupHandler defines the interface for user-group-related HTTP handlers.
// @title UserGroupHandler Interface
// @description Interface for handling user-group-related HTTP requests.
type UserGroupHandler interface {
	AddUserToGroup(c *gin.Context)
	RemoveUserFromGroup(c *gin.Context)
	GetUserGroups(c *gin.Context)
	GetGroupUsers(c *gin.Context)
}

// UserGroupImplementation handles HTTP requests for operations against the user's groups.
type UserGroupImplementation struct {
	userGroupService services.UserGroupService
}

// NewGroupHandler creates a new instance of the UserGroupImplementation.
func NewUserGroupHandler(userGroupService services.UserGroupService) *UserGroupImplementation {
	return &UserGroupImplementation{
		userGroupService: userGroupService,
	}
}

// AddUserToGroup adds a user to a group.
// @Summary Add a user to a group
// @Description Add a user to a specified group by their IDs
// @Tags user_group
// @Security BearerAuth
// @Param userId path int true "User ID"
// @Param groupId path int true "Group ID"
// @Success 204
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users-groups/{groupId}/users/{userId} [post]
func (handler *UserGroupImplementation) AddUserToGroup(c *gin.Context) {
	userId := c.Param("userId")
	groupId := c.Param("groupId")

	// Convert userId and groupId from string to uint
	uid, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	gid, err := strconv.ParseUint(groupId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	if err := handler.userGroupService.AddUserToGroup(uint(uid), uint(gid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// RemoveUserFromGroup removes a user from a group.
// @Summary Remove a user from a group
// @Description Remove a user from a specified group by their IDs
// @Tags user_group
// @Security BearerAuth
// @Param userId path int true "User ID"
// @Param groupId path int true "Group ID"
// @Success 204
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users-groups/{groupId}/users/{userId} [delete]
func (handler *UserGroupImplementation) RemoveUserFromGroup(c *gin.Context) {
	userId := c.Param("userId")
	groupId := c.Param("groupId")

	// Convert userId and groupId from string to uint
	uid, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	gid, err := strconv.ParseUint(groupId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	if err := handler.userGroupService.RemoveUserFromGroup(uint(uid), uint(gid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetUserGroups retrieves all groups for a user.
// @Summary Get all groups for a user
// @Description Get a list of all groups that a user belongs to by their ID
// @Tags user_group
// @Produce json
// @Security BearerAuth
// @Param userId path int true "User ID"
// @Success 200 {array} models.Group
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users-groups/users/{userId} [get]
func (handler *UserGroupImplementation) GetUserGroups(c *gin.Context) {
	userId := c.Param("userId")

	// Convert userId from string to uint
	uid, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	groups, err := handler.userGroupService.GetUserGroups(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GetGroupUsers retrieves all users for a group.
// @Summary Get all users for a group
// @Description Get a list of all users that belong to a group by its ID
// @Tags user_group
// @Produce json
// @Security BearerAuth
// @Param groupId path int true "Group ID"
// @Success 200 {array} models.User
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /users-groups/{groupId}/users [get]
func (handler *UserGroupImplementation) GetGroupUsers(c *gin.Context) {
	groupId := c.Param("groupId")

	// Convert groupId from string to uint
	gid, err := strconv.ParseUint(groupId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	users, err := handler.userGroupService.GetGroupUsers(uint(gid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
