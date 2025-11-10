package repository

import (
	"ceilo-backend/internal/models"

	"gorm.io/gorm"
)

// ForumRepository handles forum data operations
type ForumRepository struct {
	db *gorm.DB
}

// NewForumRepository creates a new forum repository
func NewForumRepository(db *gorm.DB) *ForumRepository {
	return &ForumRepository{db: db}
}

// Create creates a new forum post
func (r *ForumRepository) Create(forum *models.Forum) error {
	return r.db.Create(forum).Error
}

// FindAll retrieves all forum posts with pagination
func (r *ForumRepository) FindAll(page, limit int) ([]models.Forum, error) {
	var forums []models.Forum
	offset := (page - 1) * limit

	err := r.db.Preload("User").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&forums).Error

	return forums, err
}

// FindByID finds a forum post by ID
func (r *ForumRepository) FindByID(id uint) (*models.Forum, error) {
	var forum models.Forum
	err := r.db.Preload("User").
		Preload("Comments.User").
		First(&forum, id).Error
	return &forum, err
}

// Update updates a forum post
func (r *ForumRepository) Update(forum *models.Forum) error {
	return r.db.Save(forum).Error
}

// Delete deletes a forum post
func (r *ForumRepository) Delete(id uint) error {
	return r.db.Delete(&models.Forum{}, id).Error
}

// IncrementViewCount increments view count of a forum post
func (r *ForumRepository) IncrementViewCount(id uint) error {
	return r.db.Model(&models.Forum{}).
		Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}