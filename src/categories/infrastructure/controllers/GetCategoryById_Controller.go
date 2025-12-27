package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/categories/application"
	"estsoftwareoficial/src/categories/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetCategoryByIdController struct {
	getCategoryById *application.GetCategoryById
}

func NewGetCategoryByIdController(getCategoryById *application.GetCategoryById) *GetCategoryByIdController {
	return &GetCategoryByIdController{getCategoryById: getCategoryById}
}

func (gc *GetCategoryByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	category, err := gc.getCategoryById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": dto.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
		},
	})
}
