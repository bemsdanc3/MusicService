package usecases

import (
	"musicService/internal/entities"
	"musicService/internal/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]entities.User, error)
	GetUserByID(id int) (entities.User, error)
	CreateUser(user entities.User) (int64, error)
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

func (u *userUsecase) GetUserByID(id int) (entities.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userUsecase) CreateUser(user entities.User) (int64, error) {
	return u.userRepo.CreateUser(user)
}
