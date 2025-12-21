package domain

import "estsoftwareoficial/src/users/domain/entities"

type UserRepository interface {
	Save(user *entities.User) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	GetByID(id int) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Update(user *entities.User) error
	Delete(id int) error
	GetTotal() (int, error)
}
