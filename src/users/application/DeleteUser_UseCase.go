package application

import (
	"errors"
	"estsoftwareoficial/src/core/cloudinary"
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

	if existingUser.ProfilePhoto != nil && *existingUser.ProfilePhoto != "" {
		if err := cloudinary.DeleteImage(*existingUser.ProfilePhoto); err != nil {
			return errors.New("error al eliminar foto de perfil de Cloudinary")
		}
	}

	return du.userRepo.Delete(id)
}
