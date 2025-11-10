package service

import (
	"ceilo-backend/internal/models"
	"ceilo-backend/internal/repository"
)

// ConsultationService handles consultation business logic
type ConsultationService struct {
	consultationRepo *repository.ConsultationRepository
}

// NewConsultationService creates a new consultation service
func NewConsultationService(consultationRepo *repository.ConsultationRepository) *ConsultationService {
	return &ConsultationService{consultationRepo: consultationRepo}
}

// CreateConsultation creates a new consultation booking
func (s *ConsultationService) CreateConsultation(consultation *models.Consultation) error {
	consultation.Status = "pending"
	return s.consultationRepo.Create(consultation)
}

// GetConsultationsByUserID retrieves consultations by user ID
func (s *ConsultationService) GetConsultationsByUserID(userID uint) ([]models.Consultation, error) {
	return s.consultationRepo.FindByUserID(userID)
}

// GetConsultationByID retrieves a consultation by ID
func (s *ConsultationService) GetConsultationByID(id uint) (*models.Consultation, error) {
	return s.consultationRepo.FindByID(id)
}

// UpdateConsultation updates a consultation
func (s *ConsultationService) UpdateConsultation(consultation *models.Consultation) error {
	return s.consultationRepo.Update(consultation)
}

// CancelConsultation cancels a consultation
func (s *ConsultationService) CancelConsultation(id uint) error {
	consultation, err := s.consultationRepo.FindByID(id)
	if err != nil {
		return err
	}

	consultation.Status = "cancelled"
	return s.consultationRepo.Update(consultation)
}