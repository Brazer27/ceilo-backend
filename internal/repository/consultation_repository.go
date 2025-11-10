package repository

import (
	"ceilo-backend/internal/models"

	"gorm.io/gorm"
)

// ConsultationRepository handles consultation data operations
type ConsultationRepository struct {
	db *gorm.DB
}

// NewConsultationRepository creates a new consultation repository
func NewConsultationRepository(db *gorm.DB) *ConsultationRepository {
	return &ConsultationRepository{db: db}
}

// Create creates a new consultation booking
func (r *ConsultationRepository) Create(consultation *models.Consultation) error {
	return r.db.Create(consultation).Error
}

// FindByUserID retrieves consultations by user ID
func (r *ConsultationRepository) FindByUserID(userID uint) ([]models.Consultation, error) {
	var consultations []models.Consultation
	err := r.db.Preload("Psychologist").
		Where("user_id = ?", userID).
		Order("scheduled_at DESC").
		Find(&consultations).Error
	return consultations, err
}

// FindByID finds a consultation by ID
func (r *ConsultationRepository) FindByID(id uint) (*models.Consultation, error) {
	var consultation models.Consultation
	err := r.db.Preload("User").
		Preload("Psychologist").
		First(&consultation, id).Error
	return &consultation, err
}

// Update updates a consultation
func (r *ConsultationRepository) Update(consultation *models.Consultation) error {
	return r.db.Save(consultation).Error
}

// Delete deletes a consultation
func (r *ConsultationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Consultation{}, id).Error
}