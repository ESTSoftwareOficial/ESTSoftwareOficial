package application

import (
	"errors"

	"estsoftwareoficial/src/modules/domain"
)

type ReorderModules struct {
	moduleRepo domain.ModuleRepository
}

func NewReorderModules(moduleRepo domain.ModuleRepository) *ReorderModules {
	return &ReorderModules{moduleRepo: moduleRepo}
}

type ModuleOrder struct {
	ID         int
	OrderIndex int
}

func (rm *ReorderModules) Execute(modules []ModuleOrder) error {
	if len(modules) == 0 {
		return errors.New("no se proporcionaron módulos para reordenar")
	}

	for _, module := range modules {
		err := rm.moduleRepo.UpdateOrderIndex(module.ID, module.OrderIndex)
		if err != nil {
			return err
		}
	}

	return nil
}
