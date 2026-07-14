package db

import (
	"github.com/BalamuruganP25/go-agent-framework/internal/models"

	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) error {
	return database.AutoMigrate(
		&models.Message{},
	)
}
