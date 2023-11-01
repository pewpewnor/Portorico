package repository

import (
	"github.com/pewpewnor/portorico/server/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(username string, password string, name string) (model.User, error) {
	user := model.User{Username: username, Password: password, Name: name}
	return user, r.DB.Create(&user).Error
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	return users, err
}