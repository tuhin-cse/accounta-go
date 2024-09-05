package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);not null" json:"email"`
	Phone     string         `gorm:"type:varchar(255);" json:"phone"`
	Address   string         `json:"address"`
	UserID    uint           `gorm:"not null;index" json:"user_id,omitempty"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
