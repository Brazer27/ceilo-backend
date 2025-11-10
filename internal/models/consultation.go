package models

import (
	"time"

	"gorm.io/gorm"
)

// Consultation represents a consultation booking
type Consultation struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	UserID         uint           `json:"user_id"`
	User           User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	PsychologistID uint           `json:"psychologist_id"`
	Psychologist   User           `gorm:"foreignKey:PsychologistID" json:"psychologist,omitempty"`
	ScheduledAt    time.Time      `json:"scheduled_at"`
	Status         string         `gorm:"default:'pending'" json:"status"` // pending, confirmed, completed, cancelled
	Notes          string         `gorm:"type:text" json:"notes"`
	Feedback       string         `gorm:"type:text" json:"feedback"`
	Rating         int            `gorm:"default:0" json:"rating"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

// ConsultationRequest represents consultation booking request
type ConsultationRequest struct {
	PsychologistID uint      `json:"psychologist_id" binding:"required"`
	ScheduledAt    time.Time `json:"scheduled_at" binding:"required"`
	Notes          string    `json:"notes"`
}