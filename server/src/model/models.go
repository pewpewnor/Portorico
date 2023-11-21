package model

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

var Tables = []string{"user", "session"}

type Base struct {
	Id        uuid.UUID    `db:"id" json:"id"`
	CreatedAt time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time    `db:"updated_at" json:"updatedAt"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"-"`
}

func (b *Base) FillBaseInsert() {
	b.Id = uuid.New()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Base) FillBaseUpdate() {
	b.UpdatedAt = time.Now()
}

func (b *Base) FillBaseDelete() {
	b.DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}
}

type User struct {
	Base
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
}

type Session struct {
	Base
	Token  string    `db:"token" json:"token"`
	UserId uuid.UUID `db:"user_id" json:"userId"`
}

type Website struct {
	Base
	Name              string          `db:"name" json:"name"`
	TemplateName      string          `db:"template_name" json:"templateName"`
	VisitorsThisMonth int32           `db:"visitors_this_month" json:"visitorsThisMonth"`
	Content           json.RawMessage `db:"content" json:"content"`
	UserId            uuid.UUID       `db:"user_id" json:"userId"`
}
