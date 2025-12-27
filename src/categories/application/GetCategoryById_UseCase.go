package application

import (
	"estsoftwareoficial/src/categories/domain"
	"estsoftwareoficial/src/categories/domain/entities"
)

type GetCategoryById struct {
	categoryRepo domain.CategoryRepository
}

func NewGetCategoryById(categoryRepo domain.CategoryRepository) *GetCategoryById {
	return &GetCategoryById{categoryRepo: categoryRepo}
}

func (gc *GetCategoryById) Execute(id int) (*entities.Category, error) {
	return gc.categoryRepo.GetByID(id)
}
