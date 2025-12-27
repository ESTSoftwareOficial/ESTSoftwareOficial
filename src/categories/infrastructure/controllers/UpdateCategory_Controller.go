package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/categories/application"
	"estsoftwareoficial/src/categories/domain/dto"
	"estsoftwareoficial/src/categories/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateCategoryController struct {
	updateCategory *application.UpdateCategory
}

func NewUpdateCategoryController(updateCategory *application.UpdateCategory) *UpdateCategoryController {
	return &UpdateCategoryController{updateCategory: updateCategory}
}

func (uc *UpdateCategoryController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := &entities.Category{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	err = uc.updateCategory.Execute(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoría actualizada exitosamente",
	})
}
