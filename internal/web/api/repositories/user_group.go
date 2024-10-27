package repositories

import (
	"github.com/Nokeni/GODS/internal/web/api/models"
	"gorm.io/gorm"
)

// UserGroupRepository defines the methods for interacting with the group data.
type UserGroupRepository interface {
	AddUserToGroup(userID uint, groupID uint) error
	RemoveUserFromGroup(userID uint, groupID uint) error
	GetUserGroups(userID uint) ([]*models.Group, error)
	GetGroupUsers(groupID uint) ([]*models.User, error)
}

// UserGroupRepository is an implementation of the UserGroupRepository using Gorm.
type UserGroupRepositoryImplementation struct {
	database *gorm.DB
}

func NewUserGroupRepository(database *gorm.DB) UserGroupRepository {
	return &UserGroupRepositoryImplementation{database: database}
}

// AddUserToGroup adds a user to a group.
func (repo *UserGroupRepositoryImplementation) AddUserToGroup(userID uint, groupID uint) error {
	user := &models.User{}
	group := &models.Group{}

	if err := repo.database.First(user, userID).Error; err != nil {
		return err
	}
	if err := repo.database.First(group, groupID).Error; err != nil {
		return err
	}

	return repo.database.Model(user).Association("Groups").Append(group)
}

// RemoveUserFromGroup removes a user from a group.
func (repo *UserGroupRepositoryImplementation) RemoveUserFromGroup(userID uint, groupID uint) error {
	user := &models.User{}
	group := &models.Group{}

	if err := repo.database.First(user, userID).Error; err != nil {
		return err
	}
	if err := repo.database.First(group, groupID).Error; err != nil {
		return err
	}

	return repo.database.Model(user).Association("Groups").Delete(group)
}

// GetUserGroups retrieves all groups for a user.
func (repo *UserGroupRepositoryImplementation) GetUserGroups(userID uint) ([]*models.Group, error) {
	var user models.User
	if err := repo.database.Preload("Groups").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Groups, nil
}

// GetGroupUsers retrieves all users for a group.
func (repo *UserGroupRepositoryImplementation) GetGroupUsers(groupID uint) ([]*models.User, error) {
	var group models.Group
	if err := repo.database.Preload("Users").First(&group, groupID).Error; err != nil {
		return nil, err
	}
	return group.Users, nil
}
