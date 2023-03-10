package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	MONGODB_URI  string `mapstructure:"MONGODB_URI"`
	MONGODB_NAME string `mapstructure:"MONGODB_NAME"`
	PORT         string `mapstructure:"PORT"`
}

func LoadConfig() (EnvVars, error) {
	err := godotenv.Load()
	if err != nil {
		return EnvVars{}, err
	}
	return EnvVars{
		MONGODB_URI:  os.Getenv("MONGODB_URI"),
		MONGODB_NAME: os.Getenv("MONGODB_NAME"),
		PORT:         os.Getenv("PORT"),
	}, nil
}
