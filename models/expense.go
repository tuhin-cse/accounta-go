package models

import (
	"gorm.io/gorm"
	"time"
)

type Expense struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	CategoryID  uint           `gorm:"not null;index" json:"category_id"`
	Amount      float64        `gorm:"not null" json:"amount"`
	Description string         `json:"description" json:"description"`
	Date        time.Time      `json:"date" json:"date"`
	ReceiptURL  string         `json:"receipt_url"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
