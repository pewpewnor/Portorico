package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/repository"
)

type handler struct {
	userRepository *repository.LiveUserRepository
}

func NewHandler(db *sqlx.DB) *handler {
	return &handler{repository.NewLiveUserRepository(db)}
}
