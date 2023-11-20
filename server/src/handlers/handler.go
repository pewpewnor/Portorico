package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/repository"
)

type Handler struct {
	userRepository *repository.LiveUserRepository
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{repository.NewLiveUserRepository(db)}
}

func (h *Handler) validateStringNotEmpty(validations map[string]string, fieldName string, fieldValue string) {
	if fieldValue == "" {
		validations[fieldName] = fieldName + " must not be empty"
	}
}
