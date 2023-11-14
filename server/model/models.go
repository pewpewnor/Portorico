package model

var Models = []any{
	&User{},
}

type User struct {
	Base
	Username string `gorm:"index" json:"username"`
	Password string `json:"password"`
}
