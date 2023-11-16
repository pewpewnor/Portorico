package repository

import (
	"github.com/charmbracelet/log"
	"github.com/pewpewnor/portorico/server/model"
	"github.com/pewpewnor/portorico/server/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(username string, password string) (model.User, error) {
	hashedPassword, err := utils.EncryptPassword(password)
	if err != nil {
		log.Errorf("server cannot hash password when creating user: %v", err)
	}

	user := model.User{Username: username, Password: hashedPassword}
	err = r.DB.Create(&user).Error
	if err != nil {
		log.Errorf("server cannot create user: %v", err)
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	if err != nil {
		log.Errorf("server cannot get all users: %v", err)
		return nil, err
	}

	return users, nil
}
