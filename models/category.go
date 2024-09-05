package models

type Category struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Type        string `gorm:"type:varchar(50);not null" json:"type"`
	Description string `json:"description"`
	UserID      uint   `gorm:"not null;index" json:"user_id,omitempty"`
}

func (c *Category) Mask() {
	c.UserID = 0
}
