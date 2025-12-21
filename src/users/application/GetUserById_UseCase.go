package application

import (
	"estsoftwareoficial/src/users/domain"
	"estsoftwareoficial/src/users/domain/entities"
)

type GetUserById struct {
	userRepo domain.UserRepository
}

func NewGetUserById(userRepo domain.UserRepository) *GetUserById {
	return &GetUserById{userRepo: userRepo}
}

func (gu *GetUserById) Execute(id int) (*entities.User, error) {
	return gu.userRepo.GetByID(id)
}
