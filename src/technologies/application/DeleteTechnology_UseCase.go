package application

import (
	"errors"

	"estsoftwareoficial/src/technologies/domain"
)

type DeleteTechnology struct {
	technologyRepo domain.TechnologyRepository
}

func NewDeleteTechnology(technologyRepo domain.TechnologyRepository) *DeleteTechnology {
	return &DeleteTechnology{technologyRepo: technologyRepo}
}

func (dt *DeleteTechnology) Execute(id int) error {
	existingTechnology, err := dt.technologyRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingTechnology == nil {
		return errors.New("tecnología no encontrada")
	}

	return dt.technologyRepo.Delete(id)
}
