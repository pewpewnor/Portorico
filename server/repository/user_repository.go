package repository

import (
	"github.com/charmbracelet/log"
	"github.com/pewpewnor/portorico/server/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(username string, password string) (model.User, error) {
	user := model.User{Username: username, Password: password}
	err := r.DB.Create(&user).Error
	if err != nil {
		log.Errorf("server cannot create user: %v\n", err)
	}
	return user, err
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	if err != nil {
		log.Errorf("server cannot get all users: %v\n", err)
	}
	return users, err
}
