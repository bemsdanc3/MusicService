package usecases

import (
	"musicService/internal/entities"
	"musicService/internal/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]entities.User, error)
}
type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: r,
	}
}

func (u *userUsecase) GetAllUsers() ([]entities.User, error) {
	return u.userRepo.GetAllUsers()
}
