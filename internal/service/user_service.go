package service

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/lithammer/shortuuid"
	"github.com/t0nyandre/go-graphql/internal/model"
)

type UserService struct {
	db    *gorm.DB
	store *redis.Client
}

func NewUserService(db *gorm.DB, store *redis.Client) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	user.ID = shortuuid.New()
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	if err := s.db.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
