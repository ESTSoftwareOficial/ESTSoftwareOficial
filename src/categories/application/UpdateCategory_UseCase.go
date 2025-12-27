package application

import (
	"errors"

	"estsoftwareoficial/src/categories/domain"
	"estsoftwareoficial/src/categories/domain/entities"
)

type UpdateCategory struct {
	categoryRepo domain.CategoryRepository
}

func NewUpdateCategory(categoryRepo domain.CategoryRepository) *UpdateCategory {
	return &UpdateCategory{categoryRepo: categoryRepo}
}

func (uc *UpdateCategory) Execute(category *entities.Category) error {
	existingCategory, err := uc.categoryRepo.GetByID(category.ID)
	if err != nil {
		return err
	}
	if existingCategory == nil {
		return errors.New("categoría no encontrada")
	}

	categoryWithSameName, _ := uc.categoryRepo.GetByName(category.Name)
	if categoryWithSameName != nil && categoryWithSameName.ID != category.ID {
		return errors.New("ya existe otra categoría con ese nombre")
	}

	return uc.categoryRepo.Update(category)
}
