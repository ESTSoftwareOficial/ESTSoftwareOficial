package domain

import "estsoftwareoficial/src/modules/domain/entities"

type ModuleRepository interface {
	Save(module *entities.Module) (*entities.Module, error)
	GetByID(id int) (*entities.Module, error)
	GetAll() ([]*entities.Module, error)
	GetByCourse(courseID int) ([]*entities.Module, error)
	Update(module *entities.Module) error
	Delete(id int) error
	UpdateOrderIndex(id int, orderIndex int) error
	GetTotalModulesByCourse(courseID int) (int, error)
}
