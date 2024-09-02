package services

import (
	"github.com/Siddheshk02/go-blog-platform/models"
	repository "github.com/Siddheshk02/go-blog-platform/repositories"
)

type PostService interface {
	CreatePost(post models.Post) error
	FetchPost(id int) (*models.Post, error)
	FetchAllPost() (*[]models.Post, error)
	UpdatePost(id int, post models.Post) error
	DeletePost(id int) error
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{postRepo: postRepo}
}

func (s *postService) CreatePost(post models.Post) error {
	return s.postRepo.CreateNewPost(&post)
}

func (s *postService) FetchPost(id int) (*models.Post, error) {
	return s.postRepo.GetPostByID(id)
}

func (s *postService) FetchAllPost() (*[]models.Post, error) {
	return s.postRepo.GetAllPosts()
}

func (s *postService) UpdatePost(id int, post models.Post) error {
	post1, err := s.postRepo.GetPostByID(id)
	if err != nil {
		return err
	}

	post1.Title = post.Title
	post1.Content = post.Content
	return s.postRepo.UpdatePost(post1)
}

func (s *postService) DeletePost(id int) error {
	return s.postRepo.DeletePost(id)
}
