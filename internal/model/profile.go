package model

type Profile struct {
	ID        string  `json:"_id,omitempty" gorm:"type:varchar(25);unique;not null"`
	FirstName *string `json:"first_name,omitempty" gorm:"type:varchar(200)"`
	LastName  *string `json:"last_name,omitempty" gorm:"type:varchar(200)"`
	Avatar    *string `json:"avatar,omitempty"`
	Posts     []*Post
}
