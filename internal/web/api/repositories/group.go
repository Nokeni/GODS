package repositories

import (
	"github.com/Nokeni/GODS/internal/web/api/models"
	"gorm.io/gorm"
)

// GroupRepository defines the methods for interacting with the group data.
type GroupRepository interface {
	Get(id uint) (*models.Group, error)
	GetByName(name string) (*models.Group, error)
	GetAll() ([]*models.Group, error)
	Create(group *models.Group) error
	Update(group *models.Group) error
	Delete(id uint) error
}

// GroupRepositoryImplementation is an implementation of the GroupRepository using Gorm.
type GroupRepositoryImplementation struct {
	database *gorm.DB
}

func NewGroupRepository(database *gorm.DB) GroupRepository {
	return &GroupRepositoryImplementation{database: database}
}

// Get retrieves a group by ID.
func (repo *GroupRepositoryImplementation) Get(id uint) (*models.Group, error) {
	var group models.Group
	if err := repo.database.Preload("Users").First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

// GetByName retrieves a group by name.
func (repo *GroupRepositoryImplementation) GetByName(name string) (*models.Group, error) {
	var group models.Group
	if err := repo.database.Where("name = ?", name).Preload("Users").First(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

// GetAll retrieves all groups.
func (repo *GroupRepositoryImplementation) GetAll() ([]*models.Group, error) {
	var groups []*models.Group
	if err := repo.database.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// Create adds a new group.
func (repo *GroupRepositoryImplementation) Create(group *models.Group) error {
	return repo.database.Create(group).Error
}

// Update modifies an existing group.
func (repo *GroupRepositoryImplementation) Update(group *models.Group) error {
	return repo.database.Save(group).Error
}

// Delete removes a group by ID.
func (repo *GroupRepositoryImplementation) Delete(id uint) error {
	return repo.database.Delete(&models.Group{}, id).Error
}
