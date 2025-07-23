package config

import (
	"fmt"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB     *gorm.DB
	Logger *Logger
}

var appConfig *AppConfig

func Initialize() (*AppConfig, error) {
	logger := NewLogger("app")
	dbConfig := NewDatabaseConfig(logger)

	db, err := dbConfig.InitializeSQLite()
	if err != nil {
		return nil, fmt.Errorf("error initializing sqlite: %v", err)
	}

	appConfig = &AppConfig{
		DB:     db,
		Logger: logger,
	}

	return appConfig, nil
}

func GetDB() *gorm.DB {
	return appConfig.DB
}

func GetLogger() *Logger {
	return appConfig.Logger
}