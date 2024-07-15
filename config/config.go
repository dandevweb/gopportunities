package config

import (
	"fmt"
	"log"
	"os"

	"github.com/dandevweb/gopportunities/docs"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	logger     *Logger
	connection string
)

func Init() error {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	connection = os.Getenv("DB_CONNECTION")
	if connection == "" {
		return fmt.Errorf("DB_CONNECTION environment variable not set")
	}

	var err error

	if connection == "mysql" {
		db, err = InitializeMySql()
		if err != nil {
			return fmt.Errorf("error initializing mysql: %v", err)
		}
	}

	if connection == "sqlite" {
		db, err = InitializeSQLite()
		if err != nil {
			return fmt.Errorf("error initializing sqlite: %v", err)
		}
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func SwaggerInit() {
	docs.SwaggerInfo.Title = "Gopportunities API"
	docs.SwaggerInfo.Description = "This is a simple API for managing job openings"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
}
