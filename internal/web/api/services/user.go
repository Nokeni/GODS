package services

import (
	"errors"

	"github.com/Nokeni/GODS/internal/web/api/models"
	"github.com/Nokeni/GODS/internal/web/api/repositories"
	"github.com/Nokeni/GODS/internal/web/common/dtos"
)

// UserService defines the methods for performing business operations on Users.
type UserService interface {
	Get(id uint) (*models.User, error)
	GetAll() ([]*models.User, error)
	Create(userDTO *dtos.CreateUserDTO) (*models.User, error)
	Update(user *models.User, userDTO *dtos.UpdateUserDTO) error
	Delete(id uint) error
}

// UserServiceImplementation is an implementation of the UserService.
type UserServiceImplementation struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImplementation{userRepository: userRepository}
}

// Get retrieves a user by ID.
func (service *UserServiceImplementation) Get(id uint) (*models.User, error) {
	return service.userRepository.Get(id)
}

// GetAll retrieves all users.
func (service *UserServiceImplementation) GetAll() ([]*models.User, error) {
	return service.userRepository.GetAll()
}

// Create adds a new user.
func (service *UserServiceImplementation) Create(userDTO *dtos.CreateUserDTO) (*models.User, error) {
	// Check if the user already exists
	user, err := service.userRepository.GetByName(userDTO.Name)
	if err == nil {
		return user, errors.New("user already exists")
	}

	// Check the provided password strength
	if err := models.ValidatePasswordStrength(userDTO.Password); err != nil {
		return nil, err
	}

	hashedPassword, err := models.HashPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}

	// Create the user model
	user = &models.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: hashedPassword,
	}

	err = service.userRepository.Create(user)

	return user, err
}

// Update modifies an existing user.
func (service *UserServiceImplementation) Update(user *models.User, userDTO *dtos.UpdateUserDTO) error {
	// Update user details depending on provided DTO fields
	if userDTO.Name != "" {
		user.Name = userDTO.Name
	}

	if userDTO.Email != "" {
		user.Email = userDTO.Email
	}

	if userDTO.Password != "" {
		if err := models.ValidatePasswordStrength(userDTO.Password); err != nil {
			return err
		}

		hashedPassword, err := models.HashPassword(userDTO.Password)
		if err != nil {
			return err
		}

		user.Password = hashedPassword
	}

	return service.userRepository.Update(user)
}

// Delete removes a user by ID.
func (service *UserServiceImplementation) Delete(id uint) error {
	return service.userRepository.Delete(id)
}
