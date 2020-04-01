package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	ID          string     `json:"_id,omitempty" gorm:"type:varchar(25);unique;not null"`
	ProfileID   string     `json:"profile_id,omitempty" gorm:"type:varchar(25);not null"`
	Title       string     `json:"title,omitempty" gorm:"type:varchar(125);not null"`
	Excerpt     *string    `json:"excerpt,omitempty"`
	Text        string     `json:"text,omitempty" gorm:"not null"`
	Image       *string    `json:"image,omitempty"`
	Slug        string     `json:"slug,omitempty" gorm:"type:varchar(125);unique;not null"`
	Published   bool       `json:"is_published,omitempty" gorm:"default:false"`
	Archived    bool       `json:"is_archived,omitempty" gorm:"default:false"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (post *Post) IsArchived() bool {
	return post.Archived
}

func (post *Post) IsUpdated() bool {
	if post.UpdatedAt != &post.CreatedAt {
		return true
	}
	return false
}

func (post *Post) IsPublished() bool {
	return post.Published
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
