package service

import (
	"ceilo-backend/internal/models"
	"ceilo-backend/internal/repository"
)

// ForumService handles forum business logic
type ForumService struct {
	forumRepo *repository.ForumRepository
}

// NewForumService creates a new forum service
func NewForumService(forumRepo *repository.ForumRepository) *ForumService {
	return &ForumService{forumRepo: forumRepo}
}

// CreateForum creates a new forum post
func (s *ForumService) CreateForum(forum *models.Forum) error {
	return s.forumRepo.Create(forum)
}

// GetAllForums retrieves all forum posts with pagination
func (s *ForumService) GetAllForums(page, limit int) ([]models.Forum, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return s.forumRepo.FindAll(page, limit)
}

// GetForumByID retrieves a forum post by ID
func (s *ForumService) GetForumByID(id uint) (*models.Forum, error) {
	// Increment view count
	s.forumRepo.IncrementViewCount(id)
	
	return s.forumRepo.FindByID(id)
}

// UpdateForum updates a forum post
func (s *ForumService) UpdateForum(forum *models.Forum) error {
	return s.forumRepo.Update(forum)
}

// DeleteForum deletes a forum post
func (s *ForumService) DeleteForum(id uint) error {
	return s.forumRepo.Delete(id)
}