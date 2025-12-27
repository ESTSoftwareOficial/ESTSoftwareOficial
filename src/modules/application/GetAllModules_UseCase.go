package application

import (
	"estsoftwareoficial/src/modules/domain"
	"estsoftwareoficial/src/modules/domain/entities"
)

type GetAllModules struct {
	moduleRepo domain.ModuleRepository
}

func NewGetAllModules(moduleRepo domain.ModuleRepository) *GetAllModules {
	return &GetAllModules{moduleRepo: moduleRepo}
}

func (gm *GetAllModules) Execute() ([]*entities.Module, error) {
	return gm.moduleRepo.GetAll()
}
