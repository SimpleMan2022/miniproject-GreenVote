package repositories

import (
	"evoting/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	FindById(id uuid.UUID) (*entities.Comment, error)
	Create(comment *entities.Comment) (*entities.Comment, error)
	Update(comment *entities.Comment) (*entities.Comment, error)
	Delete(comment *entities.Comment) (*entities.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepositories(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) FindById(id uuid.UUID) (*entities.Comment, error) {
	var comment entities.Comment
	if err := r.db.Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) Create(comment *entities.Comment) (*entities.Comment, error) {
	if err := r.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentRepository) Update(comment *entities.Comment) (*entities.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (r *commentRepository) Delete(comment *entities.Comment) (*entities.Comment, error) {
	//TODO implement me
	panic("implement me")
}
