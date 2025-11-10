package handler

import (
	"ceilo-backend/internal/models"
	"ceilo-backend/internal/service"
	"ceilo-backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ForumHandler handles forum requests
type ForumHandler struct {
	forumService *service.ForumService
}

// NewForumHandler creates a new forum handler
func NewForumHandler(forumService *service.ForumService) *ForumHandler {
	return &ForumHandler{forumService: forumService}
}

// CreateForum creates a new forum post
func (h *ForumHandler) CreateForum(c *gin.Context) {
	var forum models.Forum
	userID := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&forum); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	forum.UserID = userID

	if err := h.forumService.CreateForum(&forum); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create forum post", err)
		return
	}

	utils.CreatedResponse(c, "Forum post created successfully", forum)
}

// GetAllForums retrieves all forum posts
func (h *ForumHandler) GetAllForums(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	forums, err := h.forumService.GetAllForums(page, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve forums", err)
		return
	}

	utils.SuccessResponse(c, "Forums retrieved successfully", forums)
}

// GetForumByID retrieves a forum post by ID
func (h *ForumHandler) GetForumByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid forum ID", err)
		return
	}

	forum, err := h.forumService.GetForumByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Forum not found", err)
		return
	}

	utils.SuccessResponse(c, "Forum retrieved successfully", forum)
}

// UpdateForum updates a forum post
func (h *ForumHandler) UpdateForum(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid forum ID", err)
		return
	}

	var forum models.Forum
	if err := c.ShouldBindJSON(&forum); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	forum.ID = uint(id)
	if err := h.forumService.UpdateForum(&forum); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update forum", err)
		return
	}

	utils.SuccessResponse(c, "Forum updated successfully", forum)
}

// DeleteForum deletes a forum post
func (h *ForumHandler) DeleteForum(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid forum ID", err)
		return
	}

	if err := h.forumService.DeleteForum(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete forum", err)
		return
	}

	utils.SuccessResponse(c, "Forum deleted successfully", nil)
}