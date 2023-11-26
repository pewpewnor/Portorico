package repository

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/model"
)

type WebsiteRepository struct {
	db *sqlx.DB
}

func NewWebsiteRepository(db *sqlx.DB) *WebsiteRepository {
	return &WebsiteRepository{db: db}
}

func (wr *WebsiteRepository) GetById(id uuid.UUID) (model.Website, bool) {
	website := model.Website{}
	if err := wr.db.Get(&website, "SELECT * FROM websites WHERE id = $1", id); err != nil {
		return model.Website{}, false
	}

	return website, true
}

func (wr *WebsiteRepository) GetByName(name string) (model.Website, bool) {
	website := model.Website{}
	if err := wr.db.Get(&website, "SELECT * FROM websites WHERE name = $1", name); err != nil {
		return model.Website{}, false
	}

	return website, true
}

func (wr *WebsiteRepository) FindByUserId(userId uuid.UUID) ([]model.Website, error) {
	websites := []model.Website{}
	err := wr.db.Select(&websites, "SELECT * FROM websites WHERE websites.user_id = $1", userId)
	if err != nil {
		log.Errorf("server cannot find websites by userId: %v", err)
		return nil, err
	}

	return websites, nil
}

func (wr *WebsiteRepository) Create(name string, templateName string, description string, userId uuid.UUID) (model.Website, error) {
	website := model.Website{Name: name, TemplateName: templateName, Description: description, VisitorsThisMonth: 0, Content: []byte("{}"), UserId: userId}
	website.FillBaseInsert()
	_, err := wr.db.NamedExec("INSERT INTO websites VALUES (:id, :created_at, :updated_at, :deleted_at, :name, :template_name, :description, :visitors_this_month, :content, :user_id)", &website)
	if err != nil {
		log.Errorf("server cannot create website: %v", err)
		return model.Website{}, err
	}

	return website, nil
}

func (wr *WebsiteRepository) Update(name string, description string, content json.RawMessage, websiteId uuid.UUID) error {
	_, err := wr.db.Exec("UPDATE websites SET name = $1, description = $2, content = $3 WHERE id = $4", name, description, content, websiteId)
	if err != nil {
		log.Errorf("server cannot update website content: %v", err)
		return err
	}

	return nil
}

func (wr *WebsiteRepository) Delete(websiteId uuid.UUID) error {
	_, err := wr.db.Exec("DELETE FROM websites WHERE id = $1", websiteId)
	if err != nil {
		log.Errorf("server cannot delete website: %v", err)
		return err
	}

	return nil
}
