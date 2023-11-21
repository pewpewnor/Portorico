package repository

import (
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/model"
	"github.com/pewpewnor/portorico/server/src/utils"
)

type LiveUserRepository struct {
	db *sqlx.DB
}

func NewLiveUserRepository(db *sqlx.DB) *LiveUserRepository {
	return &LiveUserRepository{db: db}
}

func (r *LiveUserRepository) createSession(userId uuid.UUID) (*model.Session, error) {
	session := &model.Session{Token: utils.GenerateRandomString(32), UserId: userId}
	session.FillBaseInsert()
	if _, err := r.db.NamedExec("INSERT INTO sessions VALUES (:id, :created_at, :updated_at, :deleted_at, :token, :user_id)", session); err != nil {
		log.Errorf("server cannot create session: %v", err)
		return nil, err
	}

	return session, nil
}

func (r *LiveUserRepository) GetAll() ([]model.User, error) {
	users := []model.User{}
	if err := r.db.Select(&users, "SELECT * FROM users"); err != nil {
		log.Errorf("server cannot get all users: %v", err)
		return nil, err
	}

	return users, nil
}

func (r *LiveUserRepository) GetByUsername(name string) *model.User {
	user := &model.User{}
	if err := r.db.Get(user, "SELECT * FROM users WHERE username=$1", name); err != nil {
		return nil
	}

	return user
}

func (r *LiveUserRepository) GetByCredentials(username string, password string) (*model.User, *model.Session, bool, error) {
	user := &model.User{}
	if err := r.db.Get(user, "SELECT * FROM users WHERE username=$1", username); err != nil {
		return nil, nil, false, nil
	}
	if !utils.VerifySamePassword(user.Password, password) {
		return nil, nil, false, nil
	}

	session, err := r.createSession(user.Id)
	if err != nil {
		return user, nil, false, err
	}

	return user, session, true, err
}

func (r *LiveUserRepository) GetBySessionToken(token string) *model.User {
	user := &model.User{}
	err := r.db.Get(user, "SELECT users.* FROM users JOIN sessions ON sessions.user_id = users.id WHERE sessions.token = $1", token)
	if err != nil {
		return nil
	}
	return user
}

func (r *LiveUserRepository) Create(username string, password string) (*model.User, *model.Session, error) {
	hashedPassword, err := utils.EncryptPassword(password)
	if err != nil {
		log.Errorf("server cannot hash password when creating user: %v", err)
		return nil, nil, err
	}

	user := &model.User{Username: username, Password: hashedPassword}
	user.FillBaseInsert()
	if _, err = r.db.NamedExec("INSERT INTO users VALUES (:id, :created_at, :updated_at, :deleted_at, :username, :password)", user); err != nil {
		log.Errorf("server cannot create user: %v", err)
		return nil, nil, err
	}

	session, err := r.createSession(user.Id)
	if err != nil {
		return user, nil, err
	}

	return user, session, nil
}
