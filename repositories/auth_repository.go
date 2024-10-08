package repository

import (
	"github.com/Siddheshk02/go-blog-platform/models"
	"github.com/jinzhu/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}
