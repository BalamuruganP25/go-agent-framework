package models

import "time"

type Message struct {
	ID        uint   `gorm:"primaryKey"`
	SessionID string `gorm:"index"`
	Role      string `gorm:"not null"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
}
