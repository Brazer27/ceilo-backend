package models

import (
	"time"

	"gorm.io/gorm"
)

// Forum represents a forum post
type Forum struct {
	ID        uint             `gorm:"primaryKey" json:"id"`
	Title     string           `gorm:"not null" json:"title"`
	Content   string           `gorm:"type:text;not null" json:"content"`
	Category  string           `json:"category"` // anxiety, depression, stress, general
	UserID    uint             `json:"user_id"`
	User      User             `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Comments  []ForumComment   `gorm:"foreignKey:ForumID" json:"comments,omitempty"`
	ViewCount int              `gorm:"default:0" json:"view_count"`
	IsAnonymous bool           `gorm:"default:false" json:"is_anonymous"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}

// ForumComment represents a comment on a forum post
type ForumComment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ForumID   uint           `json:"forum_id"`
	UserID    uint           `json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}