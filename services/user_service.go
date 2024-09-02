package services

import (
	"github.com/Siddheshk02/go-blog-platform/models"
	repository "github.com/Siddheshk02/go-blog-platform/repositories"
)

type UserService interface {
	FetchUser(id int) (*models.UpdateProfileInput, error)
	FetchAllUser() (*[]models.UpdateProfileInput, error)
	UpdateProfile(id int, user models.UpdateProfileInput) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) FetchUser(id int) (*models.UpdateProfileInput, error) {
	res, err := s.userRepo.GetUserByID(id)

	var user models.UpdateProfileInput

	user.Username = res.Username
	user.Email = res.Email
	user.Id = res.ID

	return &user, err
}

func (s *userService) FetchAllUser() (*[]models.UpdateProfileInput, error) {
	res, err := s.userRepo.GetAllUsers()

	users := make([]models.UpdateProfileInput, 0, len(*res))

	for _, res1 := range *res {
		var user models.UpdateProfileInput
		user.Id = res1.ID
		user.Email = res1.Email
		user.Username = res1.Username

		users = append(users, user)
	}

	return &users, err
}

func (s *userService) UpdateProfile(id int, user models.UpdateProfileInput) error {
	user1, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	user1.Email = user.Email
	user1.Username = user.Username
	return s.userRepo.UpdateUserProfile(user1)
}
