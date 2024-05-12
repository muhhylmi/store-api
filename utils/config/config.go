package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() *Configurations {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err, ", environtment will get from os env")
	}
	return &Configurations{
		DB_URI:              os.Getenv("DB_URI"),
		HOST:                os.Getenv("HOST"),
		PORT:                os.Getenv("PORT"),
		BASIC_AUTH_USERNAME: os.Getenv("BASIC_AUTH_USERNAME"),
		BASIC_AUTH_PASSWORD: os.Getenv("BASIC_AUTH_PASSWORD"),
		JWT_SECRET_KEY:      os.Getenv("JWT_SECRET_KEY"),
		API_KEY:             os.Getenv("API_KEY"),
	}
}
