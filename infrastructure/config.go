package infrastructure

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Environment  string
	App          string
	Host         string
	Port         string
	User         string
	Password     string
	Database     string
	AccessSecret string
}

func NewConfig() Config {
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalln("Error loading env file")
		}
	}

	return Config{
		Environment:  os.Getenv("ENVIRONMENT"),
		App:          os.Getenv("APP"),
		Host:         os.Getenv("HOST_POSTGRES"),
		Port:         os.Getenv("PORT_POSTGRES"),
		User:         os.Getenv("USER_POSTGRES"),
		Password:     os.Getenv("PASSWORD_POSTGRES"),
		Database:     os.Getenv("DATABASE_POSTGRES"),
		AccessSecret: os.Getenv("ACCESS_SECRET"),
	}
}
