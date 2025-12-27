package application

import (
	"estsoftwareoficial/src/modules/domain"
	"estsoftwareoficial/src/modules/domain/entities"
)

type GetModuleById struct {
	moduleRepo domain.ModuleRepository
}

func NewGetModuleById(moduleRepo domain.ModuleRepository) *GetModuleById {
	return &GetModuleById{moduleRepo: moduleRepo}
}

func (gm *GetModuleById) Execute(id int) (*entities.Module, error) {
	return gm.moduleRepo.GetByID(id)
}
