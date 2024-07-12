package config

import (
	"os"

	"github.com/dandevweb/gopportunities/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Infof("Database file does not exist, creating new database at %s", dbPath)
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing sqlite database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schema.Opening{})
	if err != nil {
		logger.Errorf("Error migrating schema: %v", err)
		return nil, err
	}

	return db, nil
}
