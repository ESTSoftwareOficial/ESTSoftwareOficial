package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"estsoftwareoficial/src/core/cloudinary"
	"estsoftwareoficial/src/courses/application"
	"estsoftwareoficial/src/courses/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateCourseController struct {
	updateCourse *application.UpdateCourse
}

func NewUpdateCourseController(updateCourse *application.UpdateCourse) *UpdateCourseController {
	return &UpdateCourseController{updateCourse: updateCourse}
}

func (uc *UpdateCourseController) Execute(c *gin.Context) {
	fmt.Println("=== INICIO UPDATE COURSE ===")

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	nameCourse := c.PostForm("nameCourse")
	description := c.PostForm("description")
	technologyID := c.PostForm("technologyId")
	categoryID := c.PostForm("categoryId")
	level := c.PostForm("level")
	durationHours := c.PostForm("durationHours")
	isActive := c.PostForm("isActive")

	if nameCourse == "" || description == "" || technologyID == "" || categoryID == "" || level == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obligatorios: nameCourse, description, technologyId, categoryId, level"})
		return
	}

	technologyIDInt, err := parseToInt(technologyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "technologyId debe ser un número"})
		return
	}

	categoryIDInt, err := parseToInt(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "categoryId debe ser un número"})
		return
	}

	var durationHoursFloat *float64
	if durationHours != "" {
		val, err := parseToFloat(durationHours)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "durationHours debe ser un número"})
			return
		}
		durationHoursFloat = &val
	}

	isActiveBool := true
	if isActive == "false" {
		isActiveBool = false
	}

	var imageURL *string
	fmt.Println("Intentando leer archivo de imagen...")
	file, header, err := c.Request.FormFile("image")
	if err == nil {
		fmt.Printf("Imagen recibida: %s, Tamaño: %d bytes\n", header.Filename, header.Size)
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("ERROR al leer bytes del archivo: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar imagen"})
			return
		}

		fmt.Println("Subiendo nueva imagen a Cloudinary...")
		uploadedURL, err := cloudinary.UploadCourseImage(fileBytes, header.Filename)
		if err != nil {
			fmt.Printf("ERROR al subir a Cloudinary: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al subir imagen"})
			return
		}

		fmt.Printf("Nueva imagen subida: %s\n", uploadedURL)
		imageURL = &uploadedURL
	} else {
		fmt.Printf("No se recibió imagen (se mantendrá la imagen actual)\n")
	}

	course := &entities.Course{
		ID:            id,
		NameCourse:    nameCourse,
		Description:   description,
		TechnologyID:  technologyIDInt,
		CategoryID:    categoryIDInt,
		Level:         level,
		ImageURL:      imageURL,
		DurationHours: durationHoursFloat,
		IsActive:      isActiveBool,
	}

	fmt.Println("Actualizando curso en BD...")
	err = uc.updateCourse.Execute(course)
	if err != nil {
		fmt.Printf("ERROR al actualizar curso: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Curso actualizado exitosamente")
	c.JSON(http.StatusOK, gin.H{
		"message": "Curso actualizado exitosamente",
	})
}
