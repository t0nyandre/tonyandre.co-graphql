package model

import (
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        string  `json:"_id,omitempty" gorm:"unique,not null"`
	FirstName string  `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     string  `json:"email,omitempty" gorm:"unique,not null"`
	Password  string  `json:"password,omitempty gorm:"not null"`
	Confirmed bool    `json:"confirmed,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (user *User) HashPassword() string {
	customParams := argon2id.Params{
		Iterations:  3,
		Memory:      4096,
		Parallelism: 1,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := argon2id.CreateHash(user.Password, &customParams)
	if err != nil {
		panic(err.Error())
	}

	return hash
}

func (user *User) VerifyPassword(password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return false
	}

	return match
}

func (user *User) IsUserConfirmed() bool {
	return user.Confirmed
}

func (user *User) IsUpdated() bool {
	uat := *user.UpdatedAt
	if uat != "" {
		return true
	}
	return false
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	t := time.Now()
	scope.SetColumn("Password", user.HashPassword())
	scope.SetColumn("CreatedAt", t.Format(time.RFC822Z))
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	t := time.Now()
	scope.SetColumn("UpdatedAt", t.Format(time.RFC822Z))
	return nil
}
