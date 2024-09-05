package models

import (
	"gorm.io/gorm"
	"time"
)

type Income struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	CategoryID  uint           `gorm:"not null;index" json:"category_id"`
	Amount      float64        `gorm:"not null" json:"amount"`
	Description string         `json:"description"`
	Date        time.Time      `json:"date"`
	Source      string         `json:"source"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
