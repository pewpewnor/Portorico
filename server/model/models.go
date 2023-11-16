package model

var Models = []any{
	&User{},
}

type User struct {
	Base
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"password"`
}
