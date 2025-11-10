package models

import (
	"time"

	"gorm.io/gorm"
)

// Article represents an educational article
type Article struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Content     string         `gorm:"type:text;not null" json:"content"`
	Category    string         `json:"category"` // tips, education, news
	Thumbnail   string         `json:"thumbnail"`
	AuthorID    uint           `json:"author_id"`
	Author      User           `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	ViewCount   int            `gorm:"default:0" json:"view_count"`
	PublishedAt time.Time      `json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}