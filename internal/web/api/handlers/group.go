package handlers

import (
	"net/http"
	"strconv"

	_ "github.com/Nokeni/GODS/internal/web/api/models"
	"github.com/Nokeni/GODS/internal/web/api/services"
	"github.com/Nokeni/GODS/internal/web/common/dtos"
	"github.com/gin-gonic/gin"
)

// GroupHandler defines the interface for user-group-related HTTP handlers.
// @title GroupHandler Interface
// @description Interface for handling user-group-related HTTP requests.
type GroupHandler interface {
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// GroupHandlerImplementation handles HTTP requests for operations against the user's groups.
type GroupHandlerImplementation struct {
	groupService services.GroupService
}

// NewGroupHandler creates a new instance of the UserGroupImplementation.
func NewGroupHandler(groupService services.GroupService) *GroupHandlerImplementation {
	return &GroupHandlerImplementation{
		groupService: groupService,
	}
}

// Get retrieves a group by ID.
// @Summary Get a group by ID
// @Description Get details of a group by its ID
// @Tags groups
// @Produce json
// @Security BearerAuth
// @Param id path int true "Group ID"
// @Success 200 {object} models.Group
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /groups/{id} [get]
func (handler *GroupHandlerImplementation) Get(c *gin.Context) {
	id := c.Param("id")

	// Convert id from string to uint
	gid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	group, err := handler.groupService.Get(uint(gid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetAll retrieves all groups.
// @Summary Get all groups
// @Description Get a list of all groups
// @Tags groups
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Group
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /groups [get]
func (handler *GroupHandlerImplementation) GetAll(c *gin.Context) {
	groups, err := handler.groupService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// Create creates a new group.
// @Summary Create a new group
// @Description Create a new group in the system
// @Tags groups
// @Accept mpfd
// @Produce json
// @Security BearerAuth
// @Param name formData string true "Group name"
// @Param description formData string false "Group description"
// @Success 200 {object} models.Group
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /groups [post]
func (handler *GroupHandlerImplementation) Create(c *gin.Context) {
	var groupDTO dtos.CreateGroupDTO
	if err := c.ShouldBind(&groupDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := handler.groupService.Create(&groupDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, group)
}

// Update updates an existing group.
// @Summary Update an existing group
// @Description Update an existing group in the system
// @Tags groups
// @Accept mpfd
// @Produce json
// @Security BearerAuth
// @Param id path int true "Group ID"
// @Param name formData string false "Group name"
// @Param description formData string false "Group description"
// @Success 200 {object} models.Group
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /groups/{id} [put]
func (handler *GroupHandlerImplementation) Update(c *gin.Context) {
	id := c.Param("id")

	// Convert id from string to uint
	gid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	var groupDTO dtos.UpdateGroupDTO
	if err := c.ShouldBind(&groupDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := handler.groupService.Get(uint(gid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := handler.groupService.Update(group, &groupDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

// Delete removes a group.
// @Summary Delete a group
// @Description Remove a group from the system
// @Tags groups
// @Security BearerAuth
// @Param id path int true "Group ID"
// @Success 204
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /groups/{id} [delete]
func (handler *GroupHandlerImplementation) Delete(c *gin.Context) {
	id := c.Param("id")

	// Convert id from string to uint
	gid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	if err := handler.groupService.Delete(uint(gid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
