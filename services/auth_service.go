package services

import (
	"github.com/Siddheshk02/go-blog-platform/models"
	repository "github.com/Siddheshk02/go-blog-platform/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(username, email, password string) error
	AuthenticateUser(email, password string) (*models.User, error)
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) RegisterUser(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.authRepo.CreateUser(user)
}

func (s *authService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
