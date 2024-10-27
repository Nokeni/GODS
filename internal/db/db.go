package db

import (
	"github.com/Nokeni/GODS/internal/web/api/models"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(viper.GetString("DB_PATH")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := database.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return database, nil
}
