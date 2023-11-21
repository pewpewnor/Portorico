package repository

import (
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

func (r *LiveWebsiteRepository) CreateWebsite(name string, templateName string, userId uuid.UUID) (*model.Website, error) {
	website := &model.Website{Name: name, TemplateName: name, VisitorsThisMonth: 0, Content: "", UserId: userId}
	website.FillBaseInsert()
	_, err := r.db.NamedExec("INSERT INTO websites VALUES (:id, :created_at, :updated_at, :deleted_at, :name, :template_name, :visitors_this_month, :content, :user_id)", website)
	if err != nil {
		log.Errorf("server cannot create website: %v", err)
		return nil, err
	}

	return website, nil
}
