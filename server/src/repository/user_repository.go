package repository

import (
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/model"
	"github.com/pewpewnor/portorico/server/src/utils"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (r *UserRepository) createSession(userId uuid.UUID) (*model.Session, error) {
	session := &model.Session{Token: utils.GenerateRandomString(32), UserId: userId}
	session.FillBaseInsert()
	if _, err := r.DB.NamedExec("INSERT INTO sessions VALUES (:id, :created_at, :updated_at, :deleted_at, :token, :user_id)", session); err != nil {
		log.Errorf("server cannot create session: %v", err)
		return nil, err
	}

	return session, nil
}

func (r *UserRepository) Login(username string, password string) (*model.User, *model.Session, bool, error) {
	var user *model.User
	if err := r.DB.Get(user, "SELECT * FROM users WHERE username=$1 LIMIT 1", username); err != nil {
		return nil, nil, false, nil
	}
	if !utils.VerifySamePassword(password, user.Password) {
		return nil, nil, false, nil
	}

	session, err := r.createSession(user.Id)
	if err != nil {
		return user, nil, false, err
	}

	return user, session, true, err
}

func (r *UserRepository) Create(username string, password string) (*model.User, *model.Session, error) {
	hashedPassword, err := utils.EncryptPassword(password)
	if err != nil {
		log.Errorf("server cannot hash password when creating user: %v", err)
		return nil, nil, err
	}

	user := &model.User{Username: username, Password: hashedPassword}
	user.FillBaseInsert()
	if _, err = r.DB.NamedExec("INSERT INTO users VALUES (:id, :created_at, :updated_at, :deleted_at, :username, :password)", user); err != nil {
		log.Errorf("server cannot create user: %v", err)
		return nil, nil, err
	}

	session, err := r.createSession(user.Id)
	if err != nil {
		return user, nil, err
	}

	return user, session, nil
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Select(&users, "SELECT * FROM users"); err != nil {
		log.Errorf("server cannot get all users: %v", err)
		return nil, err
	}

	return users, nil
}
