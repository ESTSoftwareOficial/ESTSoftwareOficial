package controllers

import (
	"net/http"

	"estsoftwareoficial/src/categories/application"
	"estsoftwareoficial/src/categories/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetAllCategoriesController struct {
	getAllCategories *application.GetAllCategories
}

func NewGetAllCategoriesController(getAllCategories *application.GetAllCategories) *GetAllCategoriesController {
	return &GetAllCategoriesController{getAllCategories: getAllCategories}
}

func (gc *GetAllCategoriesController) Execute(c *gin.Context) {
	categories, err := gc.getAllCategories.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var categoryResponses []dto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categoryResponses,
	})
}
