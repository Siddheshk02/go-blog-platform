package repository

import (
	"github.com/Siddheshk02/go-blog-platform/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	UpdateUserProfile(user *models.User) error
	GetAllUsers() (*[]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UpdateUserProfile(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error

	return &user, err
}

func (r *userRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return &users, err
}
