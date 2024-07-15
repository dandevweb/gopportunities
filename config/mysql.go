package config

import (
	"os"

	"github.com/dandevweb/gopportunities/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var host string
var port string
var user string
var password string
var database string

func InitializeMySql() (*gorm.DB, error) {
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	database = os.Getenv("DB_DATABASE")

	conn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing mysql database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&model.Opening{}, &model.User{})
	if err != nil {
		logger.Errorf("Error migrating model: %v", err)
		return nil, err
	}

	return db, nil
}
