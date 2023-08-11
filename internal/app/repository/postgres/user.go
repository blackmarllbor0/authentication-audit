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
	var user models.User
	if err := r.db.Where("id = ?", ID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r UserRepository) Block(userID uint) error {
	result := r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("blocked", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r UserRepository) IncrementFailedLoginAttempts(userID uint) (int, error) {
	result := r.db.Model(&models.User{}).
		Where("id = ?", userID).
		UpdateColumn("failed_login_attempts", gorm.Expr("failed_login_attempts + ?", 1))

	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	user, err := r.GetById(userID)
	if err != nil {
		return 0, err
	}

	return user.FailedLoginAttempts, nil
}
