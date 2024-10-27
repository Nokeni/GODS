package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading configuration file: %v", err)
	}

	return nil
}
