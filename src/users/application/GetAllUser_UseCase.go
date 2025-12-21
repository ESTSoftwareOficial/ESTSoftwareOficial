package application

import (
	"estsoftwareoficial/src/users/domain"
	"estsoftwareoficial/src/users/domain/entities"
)

type GetAllUsers struct {
	userRepo domain.UserRepository
}

func NewGetAllUsers(userRepo domain.UserRepository) *GetAllUsers {
	return &GetAllUsers{userRepo: userRepo}
}

func (gu *GetAllUsers) Execute() ([]*entities.User, error) {
	return gu.userRepo.GetAll()
}
