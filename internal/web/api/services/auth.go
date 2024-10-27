package services

import (
	"errors"
	"time"

	"github.com/Nokeni/GODS/internal/web/api/models"
	"github.com/Nokeni/GODS/internal/web/api/repositories"
	"github.com/Nokeni/GODS/internal/web/common/dtos"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// AuthService defines the methods for performing business operations on User's authentication.
type AuthService interface {
	Login(loginDTO *dtos.LoginDTO) (string, error)
	Signup(signupDTO *dtos.SignupDTO) error
}

// AuthServiceImplementation is an implementation of the UserService.
type AuthServiceImplementation struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &AuthServiceImplementation{userRepository: userRepository}
}

// Login authenticates a user.
func (service *AuthServiceImplementation) Login(loginDTO *dtos.LoginDTO) (string, error) {
	user, err := service.userRepository.GetByName(loginDTO.Name)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password)) != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := generateJWTToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Signup creates a new user.
func (service *AuthServiceImplementation) Signup(signupDTO *dtos.SignupDTO) error {
	// Check if passwords match
	if signupDTO.Password != signupDTO.PasswordConfirmation {
		return errors.New("passwords doesn't match")
	}

	// Check if the user already exists
	if _, err := service.userRepository.GetByName(signupDTO.Name); err == nil {
		return errors.New("user already exists")
	}

	// Check the provided password strength
	if err := models.ValidatePasswordStrength(signupDTO.Password); err != nil {
		return err
	}

	hashedPassword, err := models.HashPassword(signupDTO.Password)
	if err != nil {
		return err
	}

	// Create the user model
	user := &models.User{
		Name:     signupDTO.Name,
		Email:    signupDTO.Email,
		Password: hashedPassword,
	}

	return service.userRepository.Create(user)
}

// generateJWTToken generates a JWT token for the user.
func generateJWTToken(user *models.User) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &struct {
		UserID uint
		jwt.StandardClaims
	}{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
