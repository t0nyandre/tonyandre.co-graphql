package model

import (
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        string     `json:"_id,omitempty" gorm:"type:varchar(25);unique;not null"`
	Username  string     `json:"username,omitempty" gorm:"type:varchar(30);unique;not null"`
	Email     string     `json:"email,omitempty" gorm:"type:varchar(200);unique;not null"`
	Password  string     `json:"password,omitempty" gorm:"not null"`
	Confirmed bool       `json:"confirmed,omitempty" gorm:"default:false"`
	Disabled  bool       `json:"disabled,omitempty" gorm:"default:false"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
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
	if uat != user.CreatedAt {
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
