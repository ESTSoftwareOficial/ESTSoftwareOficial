package application

import (
	"errors"

	"estsoftwareoficial/src/categories/domain"
)

type DeleteCategory struct {
	categoryRepo domain.CategoryRepository
}

func NewDeleteCategory(categoryRepo domain.CategoryRepository) *DeleteCategory {
	return &DeleteCategory{categoryRepo: categoryRepo}
}

func (dc *DeleteCategory) Execute(id int) error {
	existingCategory, err := dc.categoryRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingCategory == nil {
		return errors.New("categoría no encontrada")
	}

	return dc.categoryRepo.Delete(id)
}
