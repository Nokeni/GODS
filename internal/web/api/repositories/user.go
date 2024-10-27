package repositories

import (
	"github.com/Nokeni/GODS/internal/web/api/models"
	"gorm.io/gorm"
)

// UserRepository defines the methods for interacting with the user data.
type UserRepository interface {
	Get(id uint) (*models.User, error)
	GetByName(name string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
}

// UserRepositoryImplementation is an implementation of the UserRepository using Gorm.
type UserRepositoryImplementation struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &UserRepositoryImplementation{database: database}
}

// Get retrieves a user by ID.
func (repo *UserRepositoryImplementation) Get(id uint) (*models.User, error) {
	var user models.User
	if err := repo.database.Preload("Groups").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByName retrieves a user by username.
func (repo *UserRepositoryImplementation) GetByName(name string) (*models.User, error) {
	var user models.User
	if err := repo.database.Where("name = ?", name).Preload("Groups").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll retrieves all users.
func (repo *UserRepositoryImplementation) GetAll() ([]*models.User, error) {
	var users []*models.User
	if err := repo.database.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Create adds a new user.
func (repo *UserRepositoryImplementation) Create(user *models.User) error {
	return repo.database.Create(user).Error
}

// Update modifies an existing user.
func (repo *UserRepositoryImplementation) Update(user *models.User) error {
	return repo.database.Save(user).Error
}

// Delete removes a user by ID.
func (repo *UserRepositoryImplementation) Delete(id uint) error {
	return repo.database.Delete(&models.User{}, id).Error
}
