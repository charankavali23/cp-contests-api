package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from the config file
func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}
}
