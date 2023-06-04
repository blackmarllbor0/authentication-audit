package postgres

import (
	models2 "aptekaaprel/internal/pkg/repository/postgres/models/audits"
	"aptekaaprel/internal/pkg/repository/postgres/models/sessions"
	"aptekaaprel/internal/pkg/repository/postgres/models/users"
	"fmt"

	"aptekaaprel/config/configGetter"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB

	configGetter configGetter.ConfigGetter
}

func NewRepository(configGetter configGetter.ConfigGetter) *Repository {
	return &Repository{configGetter: configGetter}
}

func (r *Repository) Connect() (*gorm.DB, error) {
	DB, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		r.configGetter.GetValueByKey("POSTGRES_HOST"),
		r.configGetter.GetValueByKey("POSTGRES_PORT"),
		r.configGetter.GetValueByKey("POSTGRES_USER"),
		r.configGetter.GetValueByKey("POSTGRES_PASSWORD"),
		r.configGetter.GetValueByKey("POSTGRES_DB"),
		r.configGetter.GetValueByKey("POSTGRES_SSL_MODE"),
	))
	if err != nil {
		return nil, err
	}

	r.db = DB

	r.autoMigrate()

	return DB, nil
}

func (r *Repository) Disconnect() error {
	return r.db.Close()
}

func (r *Repository) autoMigrate() {
	r.db.AutoMigrate(
		&users.Users{},
		&sessions.Session{},
		&models2.AuthenticationAudit{},
	)
}
