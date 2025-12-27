package application

import (
	"estsoftwareoficial/src/categories/domain"
	"estsoftwareoficial/src/categories/domain/entities"
)

type GetAllCategories struct {
	categoryRepo domain.CategoryRepository
}

func NewGetAllCategories(categoryRepo domain.CategoryRepository) *GetAllCategories {
	return &GetAllCategories{categoryRepo: categoryRepo}
}

func (gc *GetAllCategories) Execute() ([]*entities.Category, error) {
	return gc.categoryRepo.GetAll()
}
