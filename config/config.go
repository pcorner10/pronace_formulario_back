package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath("config/") // optionally look for config in the working directory

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetConfigType("yml") // REQUIRED if the config file does not have the extension in the name

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}
