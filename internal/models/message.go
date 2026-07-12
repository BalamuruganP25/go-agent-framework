package models

import "time"

type Message struct {
	ID        uint      `gorm:"primaryKey"`
	SessionID string    `gorm:"index"`
	Role      string
	Content   string
	CreatedAt time.Time
}
