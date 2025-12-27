package controllers

import (
	"net/http"

	"estsoftwareoficial/src/categories/application"
	"estsoftwareoficial/src/categories/domain/dto"
	"estsoftwareoficial/src/categories/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateCategoryController struct {
	createCategory *application.CreateCategory
}

func NewCreateCategoryController(createCategory *application.CreateCategory) *CreateCategoryController {
	return &CreateCategoryController{createCategory: createCategory}
}

func (cc *CreateCategoryController) Execute(c *gin.Context) {
	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := &entities.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	savedCategory, err := cc.createCategory.Execute(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Categoría creada exitosamente",
		"category": dto.CategoryResponse{
			ID:          savedCategory.ID,
			Name:        savedCategory.Name,
			Description: savedCategory.Description,
			CreatedAt:   savedCategory.CreatedAt,
		},
	})
}
