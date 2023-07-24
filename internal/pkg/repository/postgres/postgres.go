package postgres

import (
	"auth_audit/config/configValueGetter"
	"auth_audit/internal/pkg/repository/postgres/models/audits"
	"auth_audit/internal/pkg/repository/postgres/models/sessions"
	"auth_audit/internal/pkg/repository/postgres/models/users"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB

	configValueGetter configValueGetter.ConfigValueGetter
}

func NewRepository(configValueGetter configValueGetter.ConfigValueGetter) *Repository {
	return &Repository{configValueGetter: configValueGetter}
}

func (r *Repository) Connect() (*gorm.DB, error) {
	DB, err := gorm.Open("postgres", r.configValueGetter.GetValueByKeys("app.db.dsn"))
	if err != nil {
		return nil, err
	}

	r.db = DB

	if err := r.autoMigrate(); err != nil {
		return nil, err
	}

	return r.db, nil
}

func (r *Repository) Disconnect() error {
	return r.db.Close()
}

func (r *Repository) autoMigrate() error {
	if err := r.db.AutoMigrate(
		&users.Users{},
		&sessions.Session{},
		&audits.AuthenticationAudit{},
	).Error; err != nil {
		return err
	}

	return nil
}
