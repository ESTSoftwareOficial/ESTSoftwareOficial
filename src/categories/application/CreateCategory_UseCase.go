package application

import (
	"errors"
	"time"

	"estsoftwareoficial/src/categories/domain"
	"estsoftwareoficial/src/categories/domain/entities"
)

type CreateCategory struct {
	categoryRepo domain.CategoryRepository
}

func NewCreateCategory(categoryRepo domain.CategoryRepository) *CreateCategory {
	return &CreateCategory{categoryRepo: categoryRepo}
}

func (cc *CreateCategory) Execute(category *entities.Category) (*entities.Category, error) {
	if category.Name == "" {
		return nil, errors.New("el nombre es obligatorio")
	}

	existingCategory, _ := cc.categoryRepo.GetByName(category.Name)
	if existingCategory != nil {
		return nil, errors.New("ya existe una categoría con ese nombre")
	}

	category.CreatedAt = time.Now()
	return cc.categoryRepo.Save(category)
}
