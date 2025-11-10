package models

import (
	"time"

	"gorm.io/gorm"
)

// Event represents a mental health event or workshop
type Event struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Category    string         `json:"category"` // workshop, webinar, support_group
	Location    string         `json:"location"`
	StartTime   time.Time      `json:"start_time"`
	EndTime     time.Time      `json:"end_time"`
	MaxParticipants int        `json:"max_participants"`
	CurrentParticipants int    `gorm:"default:0" json:"current_participants"`
	OrganizerID uint           `json:"organizer_id"`
	Organizer   User           `gorm:"foreignKey:OrganizerID" json:"organizer,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}