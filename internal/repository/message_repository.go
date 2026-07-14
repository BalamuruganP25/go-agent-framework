package repository

import (
	"context"

	"github.com/BalamuruganP25/go-agent-framework/internal/models"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Save(ctx context.Context, message *models.Message) error
	GetBySession(ctx context.Context, sessionID string) ([]models.Message, error)
	DeleteSession(ctx context.Context, sessionID string) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{
		db: db,
	}
}

func (r *messageRepository) Save(ctx context.Context, message *models.Message) error {

	return r.db.
		WithContext(ctx).
		Create(message).
		Error
}

func (r *messageRepository) GetBySession(ctx context.Context, sessionID string) ([]models.Message, error) {

	var messages []models.Message

	err := r.db.
		WithContext(ctx).
		Where("session_id = ?", sessionID).
		Order("created_at asc").
		Find(&messages).
		Error

	return messages, err
}

func (r *messageRepository) DeleteSession(ctx context.Context, sessionID string) error {

	return r.db.
		WithContext(ctx).
		Where("session_id = ?", sessionID).
		Delete(&models.Message{}).
		Error
}
