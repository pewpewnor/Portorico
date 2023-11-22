package repository

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/model"
)

type LiveWebsiteRepository struct {
	db *sqlx.DB
}

func NewLiveWebsiteRepository(db *sqlx.DB) *LiveWebsiteRepository {
	return &LiveWebsiteRepository{db: db}
}

func (r *LiveWebsiteRepository) GetById(id uuid.UUID) *model.Website {
	website := &model.Website{}
	if err := r.db.Get(website, "SELECT * FROM websites WHERE id = $1", id); err != nil {
		return nil
	}

	return website
}

func (r *LiveWebsiteRepository) GetByName(name string) *model.Website {
	website := &model.Website{}
	if err := r.db.Get(website, "SELECT * FROM websites WHERE name = $1", name); err != nil {
		return nil
	}

	return website
}

func (r *LiveWebsiteRepository) FindByUserId(userId uuid.UUID) ([]model.Website, error) {
	websites := []model.Website{}
	err := r.db.Select(&websites, "SELECT * FROM websites WHERE websites.user_id = $1", userId)
	if err != nil {
		log.Errorf("server cannot find websites by userId: %v", err)
		return nil, err
	}

	return websites, nil
}

func (r *LiveWebsiteRepository) Create(name string, templateName string, description string, userId uuid.UUID) (*model.Website, error) {
	website := &model.Website{Name: name, TemplateName: templateName, Description: description, VisitorsThisMonth: 0, Content: []byte("{}"), UserId: userId}
	website.FillBaseInsert()
	_, err := r.db.NamedExec("INSERT INTO websites VALUES (:id, :created_at, :updated_at, :deleted_at, :name, :template_name, :description, :visitors_this_month, :content, :user_id)", website)
	if err != nil {
		log.Errorf("server cannot create website: %v", err)
		return nil, err
	}

	return website, nil
}

func (r *LiveWebsiteRepository) Update(name string, description string, content json.RawMessage, websiteId uuid.UUID) error {
	_, err := r.db.Exec("UPDATE websites SET name = $1, description = $2, content = $3 WHERE id = $4", name, description, content, websiteId)
	if err != nil {
		log.Errorf("server cannot update website content: %v", err)
		return err
	}

	return nil
}

func (r *LiveWebsiteRepository) Delete(websiteId uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM websites WHERE id = $1", websiteId)
	if err != nil {
		log.Errorf("server cannot delete website: %v", err)
		return err
	}

	return nil
}
