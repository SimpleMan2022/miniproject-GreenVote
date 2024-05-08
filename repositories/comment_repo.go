package repositories

import (
	"evoting/dto"
	"evoting/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	FindById(id uuid.UUID) (*entities.Comment, error)
	GetDetailPlace(id uuid.UUID) (*dto.CommentDetail, error)
	FindByPlaceId(id uuid.UUID) (*[]dto.CommentData, error)
	Create(comment *entities.Comment) (*entities.Comment, error)
	Update(comment *entities.Comment) (*entities.Comment, error)
	Delete(comment *entities.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) FindById(id uuid.UUID) (*entities.Comment, error) {
	var comment entities.Comment
	if err := r.db.Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) GetDetailPlace(id uuid.UUID) (*dto.CommentDetail, error) {
	var placeDetail dto.CommentDetail
	if err := r.db.
		Table("places").
		Select("places.name as place_name ,place_addresses.province,place_addresses.city, place_addresses.sub_district, place_addresses.street_name").
		Joins("INNER JOIN place_addresses ON place_addresses.place_id = places.id").
		Where("places.id = ?", id).
		Scan(&placeDetail).
		Error; err != nil {
		return nil, err
	}
	return &placeDetail, nil
}

func (r *commentRepository) FindByPlaceId(id uuid.UUID) (*[]dto.CommentData, error) {
	var comments []dto.CommentData
	if err := r.db.
		Table("comments").
		Select("comments.id as comment_id, users.fullname, comments.body, comments.created_at,comments.updated_at").
		Joins("INNER JOIN users ON comments.user_id = users.id").
		Where("comments.place_id = ?", id).
		Scan(&comments).
		Error; err != nil {
		return nil, err
	}
	return &comments, nil
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

func (r *commentRepository) Delete(comment *entities.Comment) error {
	if err := r.db.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
