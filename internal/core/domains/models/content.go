package models

import "time"

type Content struct {
	Id          int64      `gorm:"id"`
	Title       string     `gorm:"title"`
	Excerpt     string     `gorm:"excerpt"`
	Description string     `gorm:"description"`
	Image       string     `gorm:"image"`
	Tag         string     `gorm:"tag"`
	Status      string     `gorm:"status"`
	CreatedById int64      `gorm:"created_by_id"`
	CategoryId  int64      `gorm:"category_id"`
	User        User       `gorm:"foreignKey:CreatedById"`
	Category    Category   `gorm:"foreignKey:CategoryId"`
	CreatedAt   time.Time  `gorm:"created_at"`
	UpdatedAt   *time.Time `gorm:"updated_at"`
}
