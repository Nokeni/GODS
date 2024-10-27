package models

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User is a model that represents a user.
type User struct {
	gorm.Model
	Name     string   `gorm:"not null;unique"`        // Name is the user's name.
	Email    string   `gorm:"not null"`               // Email is the user's email.
	Password string   `gorm:"not null"`               // Password is the user's password.
	Groups   []*Group `gorm:"many2many:user_groups;"` // Groups is the list of groups the user belongs to.
}

// ValidatePasswordStrength checks if the password meets the required strength criteria using regex.
func ValidatePasswordStrength(password string) error {
	// Check for at least one uppercase letter
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	// Check for at least one lowercase letter
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}
	// Check for at least one digit
	if !regexp.MustCompile(`\d`).MatchString(password) {
		return errors.New("password must contain at least one digit")
	}
	// Check for at least one special character
	if !regexp.MustCompile(`[@$!%*?&]`).MatchString(password) {
		return errors.New("password must contain at least one special character")
	}
	// Check for minimum length of 8
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

// HashPassword hashes the provided password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
