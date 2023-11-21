package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/repository"
)

type Handler struct {
	userRepository    *repository.LiveUserRepository
	websiteRepository *repository.LiveWebsiteRepository
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{repository.NewLiveUserRepository(db), repository.NewLiveWebsiteRepository(db)}
}

func (h *Handler) validateStringNotEmpty(validations map[string]string, fieldName string, showCaseName string, fieldValue string) {
	if fieldValue == "" {
		validations[fieldName] = fieldName + " must not be empty"
	}
}
