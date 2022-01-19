package product

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id          int    `gorm:"not null;uniqueIndex;primary_key"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text;not null"`
	Ingredient  string `gorm:"type:text;not null"`
	Price       uint
	Rate        float32 `gorm:"size:255;not null"`
	Type        string  `gorm:"size:255;not null"`
	Image       string  `gorm:"size:255;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
