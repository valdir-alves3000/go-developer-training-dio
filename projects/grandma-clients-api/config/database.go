package config

import (
	"os"
	"path/filepath"

	"github.com/valdir-alves3000/go-developer-training-dio/projets/grandma-clients-api/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	dbDirectory = "./db"
	dbFile      = "grandma_clients.db"
)

type DatabaseConfig struct {
	logger *Logger
}

func NewDatabaseConfig(logger *Logger) *DatabaseConfig {
	return &DatabaseConfig{logger: logger}
}

func (dc *DatabaseConfig) InitializeSQLite() (*gorm.DB, error) {
	dbPath := filepath.Join(dbDirectory, dbFile)

	if err := dc.ensureDatabaseFileExists(dbPath); err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		dc.logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	if err := dc.migrateDatabase(db); err != nil {
		return nil, err
	}

	return db, nil
}

func (dc *DatabaseConfig) ensureDatabaseFileExists(dbPath string) error {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		dc.logger.Info("database file not found, creating...")

		if err := os.MkdirAll(dbDirectory, os.ModePerm); err != nil {
			return err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func (dc *DatabaseConfig) migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		dc.logger.Errorf("sqlite automigration error: %v", err)
		return err
	}
	return nil
}
