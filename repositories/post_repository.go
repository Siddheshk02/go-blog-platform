package repository

import (
	"github.com/Siddheshk02/go-blog-platform/models"
	"github.com/jinzhu/gorm"
)

type PostRepository interface {
	CreateNewPost(post *models.Post) error
	GetPostByID(id int) (*models.Post, error)
	GetAllPosts() (*[]models.Post, error)
	UpdatePost(user *models.Post) error
	DeletePost(id int) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) CreateNewPost(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) UpdatePost(user *models.Post) error {
	return r.db.Save(user).Error
}

func (r *postRepository) GetAllPosts() (*[]models.Post, error) {
	var posts []models.Post
	err := r.db.Find(&posts).Error
	return &posts, err
}

func (r *postRepository) GetPostByID(id int) (*models.Post, error) {
	var post models.Post
	err := r.db.Where("id = ?", id).First(&post).Error

	return &post, err
}

func (r *postRepository) DeletePost(id int) error {
	return r.db.Delete(&models.Post{}, id).Error
}
