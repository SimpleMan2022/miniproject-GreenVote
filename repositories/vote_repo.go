package repositories

import (
	"evoting/dto"
	"evoting/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoteRepository interface {
	TotalVoters() (int64, error)
	TotalVotesReceived() (int64, error)
	FindUserById(id uuid.UUID) (*entities.Vote, error)
	GetTotalVotes() (*[]dto.GetPlaceWithTotalVotes, error)
	Create(vote *entities.Vote) (*entities.Vote, error)
}

type voteRepository struct {
	db *gorm.DB
}

func NewvoteRepository(db *gorm.DB) *voteRepository {
	return &voteRepository{db}
}

func (r *voteRepository) FindUserById(id uuid.UUID) (*entities.Vote, error) {
	var user entities.Vote
	if err := r.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *voteRepository) GetTotalVotes() (*[]dto.GetPlaceWithTotalVotes, error) {
	var votes *[]dto.GetPlaceWithTotalVotes

	err := r.db.Table("places").
		Select("places.id as place_id, places.name as place_name, COUNT(votes.id) as total_vote," +
			"ROUND(COUNT(votes.id) * 100.0 / (SELECT COUNT(DISTINCT user_id) FROM votes), 2) as percentage").
		Joins("LEFT JOIN votes ON places.id = votes.place_id").
		Group("places.id").
		Scan(&votes).Error
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (r *voteRepository) TotalVotesReceived() (int64, error) {
	var totalVotes int64
	err := r.db.Model(&entities.Vote{}).Count(&totalVotes).Error
	if err != nil {
		return 0, err
	}
	return totalVotes, nil
}

func (r *voteRepository) TotalVoters() (int64, error) {
	var totalVotes int64
	err := r.db.Model(&entities.User{}).Count(&totalVotes).Error
	if err != nil {
		return 0, err
	}
	return totalVotes, nil
}

func (r *voteRepository) Create(vote *entities.Vote) (*entities.Vote, error) {
	if err := r.db.Create(&vote).Error; err != nil {
		return nil, err
	}
	return vote, nil
}
