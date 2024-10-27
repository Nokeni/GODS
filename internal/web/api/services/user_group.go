package services

import (
	"github.com/Nokeni/GODS/internal/web/api/models"
	"github.com/Nokeni/GODS/internal/web/api/repositories"
)

// UserGroupService defines the methods for performing business operations on Groups.
type UserGroupService interface {
	AddUserToGroup(userID uint, groupID uint) error
	RemoveUserFromGroup(userID uint, groupID uint) error
	GetUserGroups(userID uint) ([]*models.Group, error)
	GetGroupUsers(groupID uint) ([]*models.User, error)
}

// UserGroupServiceImplementation is an implementation of the GroupService.
type UserGroupServiceImplementation struct {
	userGroupRepository repositories.UserGroupRepository
}

func NewUserGroupService(userGroupRepository repositories.UserGroupRepository) UserGroupService {
	return &UserGroupServiceImplementation{userGroupRepository: userGroupRepository}
}

// AddUserToGroup adds a user to a group.
func (service *UserGroupServiceImplementation) AddUserToGroup(userID uint, groupID uint) error {
	return service.userGroupRepository.AddUserToGroup(userID, groupID)
}

// RemoveUserFromGroup removes a user from a group.
func (service *UserGroupServiceImplementation) RemoveUserFromGroup(userID uint, groupID uint) error {
	return service.userGroupRepository.RemoveUserFromGroup(userID, groupID)
}

// GetUserGroups retrieves all groups for a user.
func (service *UserGroupServiceImplementation) GetUserGroups(userID uint) ([]*models.Group, error) {
	return service.userGroupRepository.GetUserGroups(userID)
}

// GetGroupUsers retrieves all users for a group.
func (service *UserGroupServiceImplementation) GetGroupUsers(groupID uint) ([]*models.User, error) {
	return service.userGroupRepository.GetGroupUsers(groupID)
}
