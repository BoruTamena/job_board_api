package services

import (
	"github.com/BoruTamena/job_board/internal/domain/models/dto"
	"github.com/BoruTamena/job_board/internal/domain/ports"
)

type UserServices struct {
	userRepo ports.UserRepository
}

func NewUserServices(userRepo ports.UserRepository) *UserServices {

	return &UserServices{
		userRepo: userRepo,
	}
}

func (uservice UserServices) CreateUser(user dto.User) error {

	return nil

}
