package service

import (
	"crypto/md5"
	"fmt"
	"time"

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
	return &UserService{
		db:    db,
		store: store,
	}
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	user.ID = shortuuid.New()
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	token := fmt.Sprintf("%x", md5.Sum([]byte(user.ID)))
	_, err := s.store.Set(fmt.Sprintf("%s%s", "ConfirmAccount:", token), user.ID, time.Second*60*60*24*3).Result()
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n%s\n", token)

	if err := s.db.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) ConfirmUser(token string) (bool, error) {
	var err error
	user := &model.User{}

	user.ID, err = s.store.Get(fmt.Sprintf("%s%s", "ConfirmAccount:", token)).Result()
	if err != nil {
		return false, err
	}

	if err = s.db.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return false, err
	}

	user.Confirmed = true
	if err = s.db.Save(user).Error; err != nil {
		return false, err
	}
	_, err = s.store.Del(fmt.Sprintf("%s%s", "ConfirmAccount:", token)).Result()
	if err != nil {
		return false, err
	}

	return user.Confirmed, nil
}

func (s *UserService) FindUser(user *model.User) (*model.User, error) {
	if err := s.db.Where(user).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
