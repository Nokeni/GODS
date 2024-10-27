package models

import (
	"gorm.io/gorm"
)

// Group is a model that represents a group of users.
type Group struct {
	gorm.Model
	Name        string  `gorm:"not null;unique"` // Name is the group's name
	Description string  // Description is the group's description
	Users       []*User `gorm:"many2many:user_groups;"` // Users is the list of users that belongs to the group
}
