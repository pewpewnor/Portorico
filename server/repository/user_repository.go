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

func (r *UserRepository) createSession(user model.User) (model.Session, error) {
	token, err := utils.GenerateRandomString(32)
	if err != nil {
		log.Errorf("server cannot generate token when creating session: %v", err)
		return model.Session{}, err
	}

	session := model.Session{UserID: user.ID, Token: token, User: user}
	if err = r.DB.Create(&session).Error; err != nil {
		log.Errorf("server cannot create session: %v", err)
		return model.Session{}, err
	}

	return session, nil
}

func (r *UserRepository) Login(username string, password string) (model.Session, bool, error) {
	var user model.User
	result := r.DB.Preload("session").Where("username = ?", username).Limit(1).Find(&user)
	if err := result.Error; err != nil {
		log.Errorf("server cannot find user: %v", err)
		return model.Session{}, false, err
	}

	if result.RowsAffected == 0 || utils.VerifySamePassword(password, user.Password) {
		return model.Session{}, false, nil
	}

	session, err := r.createSession(user)
	if err != nil {
		return model.Session{}, false, err
	}

	return session, true, nil
}

func (r *UserRepository) Create(username string, password string) (model.Session, error) {
	hashedPassword, err := utils.EncryptPassword(password)
	if err != nil {
		log.Errorf("server cannot hash password when creating user: %v", err)
		return model.Session{}, err
	}

	user := model.User{Username: username, Password: hashedPassword}
	if err = r.DB.Create(&user).Error; err != nil {
		log.Errorf("server cannot create user: %v", err)
		return model.Session{}, err
	}

	session, err := r.createSession(user)
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
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
