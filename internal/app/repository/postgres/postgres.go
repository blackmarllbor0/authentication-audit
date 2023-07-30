package postgres

import (
	"auth_audit/config/configValueGetter"
	"auth_audit/internal/app/repository/models"
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

	//if err := DB.DB().Ping(); err != nil {
	//	return nil, err
	//}

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
		&models.Session{},
		&models.AuthenticationAudit{},
		&models.User{},
	).Error; err != nil {
		return err
	}

	return nil
}
