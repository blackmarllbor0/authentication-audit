package postgres

import (
	"auth_audit/internal/app/repository/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) Create(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r UserRepository) GetByLogin(login string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("login = ?", login).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r UserRepository) GetById(ID uint) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r UserRepository) Block() {
	//TODO implement me
	panic("implement me")
}
