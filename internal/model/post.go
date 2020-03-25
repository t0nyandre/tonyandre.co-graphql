package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	ID          string  `json:"_id,omitempty" gorm:"unique,not null"`
	Title       string  `json:"title,omitempty"`
	Excerpt     *string `json:"excerpt,omitempty"`
	Text        string  `json:"text,omitempty"`
	Image       *string `json:"image,omitempty"`
	Slug        string  `json:"slug,omitempty" gorm:"unique,not null"`
	Archived    bool    `json:"is_archived,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
	PublishedAt *string `json:"published_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

func (post *Post) IsArchived() bool {
	return post.Archived
}

func (post *Post) IsUpdated() bool {
	uat := *post.UpdatedAt
	if uat != "" {
		return true
	}
	return false
}

func (post *Post) IsPublished() bool {
	if post.PublishedAt != nil {
		return true
	}
	return false
}

func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	t := time.Now()
	scope.SetColumn("CreatedAt", t.Format(time.RFC822Z))
	scope.SetColumn("Archived", false)
	return nil
}

func (post *Post) BeforeUpdate(scope *gorm.Scope) error {
	t := time.Now()
	scope.SetColumn("UpdatedAt", t.Format(time.RFC822Z))
	return nil
}
