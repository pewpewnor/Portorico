package model

import "github.com/google/uuid"

var Models = []any{
	&User{},
}

type Session struct {
	UserID uuid.UUID `gorm:"type:uuid"`
	Token  string    `gorm:"size:32"`
	User   User
}

type User struct {
	Base
	Username string `gorm:"uniqueIndex"`
	Password string
}
