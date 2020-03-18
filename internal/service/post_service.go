package service

import (
	"github.com/jinzhu/gorm"
	"github.com/lithammer/shortuuid"
	"github.com/t0nyandre/go-graphql/internal/model"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) CreatePost(post *model.Post) (*model.Post, error) {
	post.ID = shortuuid.New()
	if err := s.db.Create(post).Error; err != nil {
		return nil, err
	}
	if err := s.db.Where("id = ?", post.ID).First(&post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) AllPosts() ([]*model.Post, error) {
	posts := make([]*model.Post, 0)
	if err := s.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetPost(id string) (*model.Post, error) {
	post := &model.Post{ID: id}
	if err := s.db.Where(post).First(&post).Error; err != nil {
		return nil, err
	}
	return post, nil
}
