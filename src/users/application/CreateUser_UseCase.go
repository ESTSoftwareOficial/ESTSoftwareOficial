package application

import (
	"errors"
	"estsoftwareoficial/src/users/domain"
	"estsoftwareoficial/src/users/domain/entities"
	"time"
)

type CreateUser struct {
	userRepo domain.UserRepository
}

func NewCreateUser(userRepo domain.UserRepository) *CreateUser {
	return &CreateUser{userRepo: userRepo}
}

func (cu *CreateUser) Execute(user *entities.User) (*entities.User, error) {
	if user.FirstName == "" {
		return nil, errors.New("el nombre es obligatorio")
	}
	if user.Email == "" {
		return nil, errors.New("el email es obligatorio")
	}
	if user.Password == nil || *user.Password == "" {
		return nil, errors.New("la contraseña es obligatoria")
	}

	existingUser, _ := cu.userRepo.GetByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("el email ya está registrado")
	}

	user.RegistrationDate = time.Now()
	return cu.userRepo.Save(user)
}
