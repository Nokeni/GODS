package services

import (
	"errors"

	"github.com/Nokeni/GODS/internal/web/api/models"
	"github.com/Nokeni/GODS/internal/web/api/repositories"
	"github.com/Nokeni/GODS/internal/web/common/dtos"
)

// GroupService defines the methods for performing business operations on Groups.
type GroupService interface {
	Get(id uint) (*models.Group, error)
	GetAll() ([]*models.Group, error)
	Create(groupDTO *dtos.CreateGroupDTO) (*models.Group, error)
	Update(group *models.Group, groupDTO *dtos.UpdateGroupDTO) error
	Delete(id uint) error
}

// GroupServiceImplementation is an implementation of the GroupService.
type GroupServiceImplementation struct {
	groupRepository repositories.GroupRepository
}

func NewGroupService(groupRepository repositories.GroupRepository) GroupService {
	return &GroupServiceImplementation{groupRepository: groupRepository}
}

// Get retrieves a group by ID.
func (service *GroupServiceImplementation) Get(id uint) (*models.Group, error) {
	return service.groupRepository.Get(id)
}

// GetAll retrieves all groups.
func (service *GroupServiceImplementation) GetAll() ([]*models.Group, error) {
	return service.groupRepository.GetAll()
}

// Create adds a new group.
func (service *GroupServiceImplementation) Create(groupDTO *dtos.CreateGroupDTO) (*models.Group, error) {
	// Check if the group already exists
	group, err := service.groupRepository.GetByName(groupDTO.Name)
	if err == nil {
		return group, errors.New("group already exists")
	}

	// Create the group model
	group = &models.Group{
		Name:        groupDTO.Name,
		Description: groupDTO.Description,
	}

	err = service.groupRepository.Create(group)

	return group, err
}

// Update modifies an existing group.
func (service *GroupServiceImplementation) Update(group *models.Group, groupDTO *dtos.UpdateGroupDTO) error {
	// Update group details depending on provided DTO fields
	if groupDTO.Name != "" {
		group.Name = groupDTO.Name
	}

	if groupDTO.Description != "" {
		group.Description = groupDTO.Description
	}

	return service.groupRepository.Update(group)
}

// Delete removes a group by ID.
func (service *GroupServiceImplementation) Delete(id uint) error {
	return service.groupRepository.Delete(id)
}
