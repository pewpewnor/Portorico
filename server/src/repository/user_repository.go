package repository

import (
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/model"
	"github.com/pewpewnor/portorico/server/src/utils"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) createSession(userId uuid.UUID) (model.Session, error) {
	session := model.Session{Token: utils.GenerateRandomString(32), UserId: userId}
	session.FillBaseInsert()
	if _, err := ur.db.NamedExec("INSERT INTO sessions VALUES (:id, :created_at, :updated_at, :deleted_at, :token, :user_id)", &session); err != nil {
		log.Errorf("server cannot create session: %v", err)
		return model.Session{}, err
	}

	return session, nil
}

func (ur *UserRepository) Find() ([]model.User, error) {
	users := []model.User{}
	if err := ur.db.Select(&users, "SELECT * FROM users"); err != nil {
		log.Errorf("server cannot get all users: %v", err)
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetByUsername(name string) (model.User, bool) {
	user := model.User{}
	if err := ur.db.Get(&user, "SELECT * FROM users WHERE username = $1", name); err != nil {
		return model.User{}, false
	}

	return user, true
}

func (ur *UserRepository) GetByCredentials(username string, password string) (model.User, model.Session, bool, error) {
	user := model.User{}
	if err := ur.db.Get(&user, "SELECT * FROM users WHERE username = $1", username); err != nil || !utils.VerifySamePassword(user.Password, password) {
		return model.User{}, model.Session{}, false, nil
	}

	session, err := ur.createSession(user.Id)
	if err != nil {
		return user, model.Session{}, false, err
	}

	return user, session, true, err
}

func (ur *UserRepository) GetBySessionToken(token string) (model.User, bool) {
	user := model.User{}
	err := ur.db.Get(&user, "SELECT users.* FROM users JOIN sessions ON sessions.user_id = users.id WHERE sessions.token = $1", token)
	if err != nil {
		return model.User{}, false
	}
	return user, true
}

func (ur *UserRepository) Create(username string, password string) (model.User, model.Session, error) {
	hashedPassword, err := utils.EncryptPassword(password)
	if err != nil {
		log.Errorf("server cannot hash password when creating user: %v", err)
		return model.User{}, model.Session{}, err
	}

	user := model.User{Username: username, Password: hashedPassword}
	user.FillBaseInsert()
	if _, err = ur.db.NamedExec("INSERT INTO users VALUES (:id, :created_at, :updated_at, :deleted_at, :username, :password)", &user); err != nil {
		log.Errorf("server cannot create user: %v", err)
		return model.User{}, model.Session{}, err
	}

	session, err := ur.createSession(user.Id)
	if err != nil {
		return user, model.Session{}, err
	}

	return user, session, nil
}
