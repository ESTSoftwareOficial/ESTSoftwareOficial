package application

import (
	"estsoftwareoficial/src/technologies/domain"
	"estsoftwareoficial/src/technologies/domain/entities"
)

type GetTechnologyById struct {
	technologyRepo domain.TechnologyRepository
}

func NewGetTechnologyById(technologyRepo domain.TechnologyRepository) *GetTechnologyById {
	return &GetTechnologyById{technologyRepo: technologyRepo}
}

func (gt *GetTechnologyById) Execute(id int) (*entities.Technology, error) {
	return gt.technologyRepo.GetByID(id)
}
