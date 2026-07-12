package db

import (
	"github.com/BalamuruganP25/go-agent-framework/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(
		sqlite.Open(dbPath),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Message{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
