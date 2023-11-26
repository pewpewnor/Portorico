package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/repository"
)

type handler struct {
	userRepo    *repository.UserRepository
	websiteRepo *repository.WebsiteRepository
}

func NewHandler(db *sqlx.DB) *handler {
	return &handler{repository.NewUserRepository(db), repository.NewWebsiteRepository(db)}
}

func (h *handler) validateStringNotEmpty(validations map[string]string, fieldName string, showcaseName string, fieldValue string) {
	if fieldValue == "" {
		validations[fieldName] = showcaseName + " must not be empty"
		return
	}
}

func (h *handler) validateStringMaxLength(validations map[string]string, fieldName string, showcaseName string, max uint16, fieldValue string) {
	if fieldValue == "" {
		validations[fieldName] = showcaseName + " must not be empty"
		return
	}
	if uint16(len(fieldValue)) > max {
		validations[fieldName] = fmt.Sprintf("%v has a maximum of %v characters", showcaseName, max)
	}
}

func (h *handler) validateStringMinMaxLength(validations map[string]string, fieldName string, showcaseName string, min uint16, max uint16, fieldValue string) {
	if fieldValue == "" {
		validations[fieldName] = showcaseName + " must not be empty"
		return
	}
	length := uint16(len(fieldValue))
	if length < min {
		validations[fieldName] = fmt.Sprintf("%v must have atleast %v characters", showcaseName, min)
		return
	}
	if length > max {
		validations[fieldName] = fmt.Sprintf("%v has a maximum of %v characters", showcaseName, max)
	}
}

func (h *handler) validateJSONString(validations map[string]string, fieldName string, jsonString string) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		validations[fieldName] = fieldName + " is not valid json"
	}
}
