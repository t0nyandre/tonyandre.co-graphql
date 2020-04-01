package service

import (
	"github.com/jinzhu/gorm"
	"github.com/lithammer/shortuuid"
	"github.com/t0nyandre/go-graphql/internal/model"
)

type ProfileService struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{
		db: db,
	}
}

func (s *ProfileService) CreateAuthor() (*model.Profile, error) {
	profile := &model.Profile{
		ID: shortuuid.New(),
	}
	if err := s.db.Create(profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}
