package models

import (
	"time"

	"gorm.io/gorm"
)

// StressTest represents a stress test result
type StressTest struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Score       int            `json:"score"`
	Level       string         `json:"level"` // low, medium, high
	Answers     string         `gorm:"type:jsonb" json:"answers"` // JSON string of answers
	Suggestion  string         `gorm:"type:text" json:"suggestion"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// StressTestRequest represents stress test submission
type StressTestRequest struct {
	Answers map[string]int `json:"answers" binding:"required"`
}