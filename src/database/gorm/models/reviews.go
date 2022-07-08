package models

import "time"

type Review struct {
	ID_Review uint      `gorm:"primaryKey" json:"id_review"`
	ProductID int       `gorm:"foreignKey" json:"product_id"`
	UserID    int       `gorm:"foreignKey" json:"user_id"`
	Comment   string    `json:"comment"`
	Rating    float32   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Reviews []Review
