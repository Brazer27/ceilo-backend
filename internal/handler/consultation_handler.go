package handler

import (
	"ceilo-backend/internal/models"
	"ceilo-backend/internal/service"
	"ceilo-backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ConsultationHandler handles consultation requests
type ConsultationHandler struct {
	consultationService *service.ConsultationService
}

// NewConsultationHandler creates a new consultation handler
func NewConsultationHandler(consultationService *service.ConsultationService) *ConsultationHandler {
	return &ConsultationHandler{consultationService: consultationService}
}

// CreateConsultation creates a new consultation booking
func (h *ConsultationHandler) CreateConsultation(c *gin.Context) {
	var req models.ConsultationRequest
	userID := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	consultation := &models.Consultation{
		UserID:         userID,
		PsychologistID: req.PsychologistID,
		ScheduledAt:    req.ScheduledAt,
		Notes:          req.Notes,
	}

	if err := h.consultationService.CreateConsultation(consultation); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create consultation", err)
		return
	}

	utils.CreatedResponse(c, "Consultation booked successfully", consultation)
}

// GetUserConsultations retrieves consultations for current user
func (h *ConsultationHandler) GetUserConsultations(c *gin.Context) {
	userID := c.GetUint("user_id")

	consultations, err := h.consultationService.GetConsultationsByUserID(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve consultations", err)
		return
	}

	utils.SuccessResponse(c, "Consultations retrieved successfully", consultations)
}

// GetConsultationByID retrieves a consultation by ID
func (h *ConsultationHandler) GetConsultationByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid consultation ID", err)
		return
	}

	consultation, err := h.consultationService.GetConsultationByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Consultation not found", err)
		return
	}

	utils.SuccessResponse(c, "Consultation retrieved successfully", consultation)
}

// UpdateConsultationStatus updates consultation status
func (h *ConsultationHandler) UpdateConsultationStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid consultation ID", err)
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	consultation, err := h.consultationService.GetConsultationByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Consultation not found", err)
		return
	}

	consultation.Status = req.Status
	if err := h.consultationService.UpdateConsultation(consultation); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update consultation", err)
		return
	}

	utils.SuccessResponse(c, "Consultation status updated successfully", consultation)
}