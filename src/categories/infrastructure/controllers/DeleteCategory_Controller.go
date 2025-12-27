package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/categories/application"

	"github.com/gin-gonic/gin"
)

type DeleteCategoryController struct {
	deleteCategory *application.DeleteCategory
}

func NewDeleteCategoryController(deleteCategory *application.DeleteCategory) *DeleteCategoryController {
	return &DeleteCategoryController{deleteCategory: deleteCategory}
}

func (dc *DeleteCategoryController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dc.deleteCategory.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoría eliminada exitosamente",
	})
}
