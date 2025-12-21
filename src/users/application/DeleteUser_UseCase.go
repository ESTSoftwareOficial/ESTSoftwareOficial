package application

import (
	"errors"
	"estsoftwareoficial/src/users/domain"
)

type DeleteUser struct {
	userRepo domain.UserRepository
}

func NewDeleteUser(userRepo domain.UserRepository) *DeleteUser {
	return &DeleteUser{userRepo: userRepo}
}

func (du *DeleteUser) Execute(id int) error {
	existingUser, err := du.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("usuario no encontrado")
	}

	return du.userRepo.Delete(id)
}
