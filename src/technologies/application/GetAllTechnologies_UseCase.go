package application

import (
	"estsoftwareoficial/src/technologies/domain"
	"estsoftwareoficial/src/technologies/domain/entities"
)

type GetAllTechnologies struct {
	technologyRepo domain.TechnologyRepository
}

func NewGetAllTechnologies(technologyRepo domain.TechnologyRepository) *GetAllTechnologies {
	return &GetAllTechnologies{technologyRepo: technologyRepo}
}

func (gt *GetAllTechnologies) Execute() ([]*entities.Technology, error) {
	return gt.technologyRepo.GetAll()
}
