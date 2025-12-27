package application

import (
	"estsoftwareoficial/src/modules/domain"
	"estsoftwareoficial/src/modules/domain/entities"
)

type GetModulesByCourse struct {
	moduleRepo domain.ModuleRepository
}

func NewGetModulesByCourse(moduleRepo domain.ModuleRepository) *GetModulesByCourse {
	return &GetModulesByCourse{moduleRepo: moduleRepo}
}

func (gm *GetModulesByCourse) Execute(courseID int) ([]*entities.Module, error) {
	return gm.moduleRepo.GetByCourse(courseID)
}
