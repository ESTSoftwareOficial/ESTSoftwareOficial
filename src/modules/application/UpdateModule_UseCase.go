package application

import (
	"errors"

	"estsoftwareoficial/src/modules/domain"
	"estsoftwareoficial/src/modules/domain/entities"
)

type UpdateModule struct {
	moduleRepo domain.ModuleRepository
}

func NewUpdateModule(moduleRepo domain.ModuleRepository) *UpdateModule {
	return &UpdateModule{moduleRepo: moduleRepo}
}

func (um *UpdateModule) Execute(module *entities.Module) error {
	existingModule, err := um.moduleRepo.GetByID(module.ID)
	if err != nil {
		return err
	}
	if existingModule == nil {
		return errors.New("módulo no encontrado")
	}

	if module.Title == "" {
		return errors.New("el título es obligatorio")
	}

	return um.moduleRepo.Update(module)
}
