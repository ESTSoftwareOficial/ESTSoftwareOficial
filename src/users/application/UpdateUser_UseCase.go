package application

import (
	"errors"
	"estsoftwareoficial/src/users/domain"
	"estsoftwareoficial/src/users/domain/entities"
)

type UpdateUser struct {
	userRepo domain.UserRepository
}

func NewUpdateUser(userRepo domain.UserRepository) *UpdateUser {
	return &UpdateUser{userRepo: userRepo}
}

func (uu *UpdateUser) Execute(user *entities.User) error {
	existingUser, err := uu.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("usuario no encontrado")
	}

	if user.ProfilePhoto == nil || *user.ProfilePhoto == "" {
		user.ProfilePhoto = existingUser.ProfilePhoto
	}

	return uu.userRepo.Update(user)
}
