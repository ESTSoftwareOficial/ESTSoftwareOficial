package controllers

import (
	"fmt"
	"io"
	"net/http"

	"estsoftwareoficial/src/core/cloudinary"
	"estsoftwareoficial/src/courses/application"
	"estsoftwareoficial/src/courses/domain/dto"
	"estsoftwareoficial/src/courses/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateCourseController struct {
	createCourse *application.CreateCourse
}

func NewCreateCourseController(createCourse *application.CreateCourse) *CreateCourseController {
	return &CreateCourseController{createCourse: createCourse}
}

func (cc *CreateCourseController) Execute(c *gin.Context) {
	fmt.Println("=== INICIO CREATE COURSE ===")

	nameCourse := c.PostForm("nameCourse")
	description := c.PostForm("description")
	technologyID := c.PostForm("technologyId")
	instructorID := c.PostForm("instructorId")
	categoryID := c.PostForm("categoryId")
	level := c.PostForm("level")
	durationHours := c.PostForm("durationHours")

	fmt.Printf("Datos recibidos - Curso: %s, Instructor: %s\n", nameCourse, instructorID)

	if nameCourse == "" || description == "" || technologyID == "" || instructorID == "" || categoryID == "" || level == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obligatorios: nameCourse, description, technologyId, instructorId, categoryId, level"})
		return
	}

	technologyIDInt, err := parseToInt(technologyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "technologyId debe ser un número"})
		return
	}

	instructorIDInt, err := parseToInt(instructorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "instructorId debe ser un número"})
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

		fmt.Println("Subiendo imagen a Cloudinary...")
		uploadedURL, err := cloudinary.UploadCourseImage(fileBytes, header.Filename)
		if err != nil {
			fmt.Printf("ERROR al subir a Cloudinary: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al subir imagen"})
			return
		}

		fmt.Printf("Imagen subida exitosamente: %s\n", uploadedURL)
		imageURL = &uploadedURL
	} else {
		fmt.Printf("No se recibió imagen: %v\n", err)
	}

	course := &entities.Course{
		NameCourse:    nameCourse,
		Description:   description,
		TechnologyID:  technologyIDInt,
		InstructorID:  instructorIDInt,
		CategoryID:    categoryIDInt,
		Level:         level,
		ImageURL:      imageURL,
		DurationHours: durationHoursFloat,
	}

	fmt.Println("Guardando curso en BD...")
	savedCourse, err := cc.createCourse.Execute(course)
	if err != nil {
		fmt.Printf("ERROR al guardar curso: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Curso creado exitosamente")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Curso creado exitosamente",
		"course": dto.CourseResponse{
			ID:            savedCourse.ID,
			NameCourse:    savedCourse.NameCourse,
			Description:   savedCourse.Description,
			TechnologyID:  savedCourse.TechnologyID,
			InstructorID:  savedCourse.InstructorID,
			CategoryID:    savedCourse.CategoryID,
			Level:         savedCourse.Level,
			ImageURL:      savedCourse.ImageURL,
			TotalModules:  savedCourse.TotalModules,
			AverageRating: savedCourse.AverageRating,
			TotalRatings:  savedCourse.TotalRatings,
			DurationHours: savedCourse.DurationHours,
			CreatedAt:     savedCourse.CreatedAt,
			UpdatedAt:     savedCourse.UpdatedAt,
			IsActive:      savedCourse.IsActive,
		},
	})
}

func parseToInt(s string) (int, error) {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
}

func parseToFloat(s string) (float64, error) {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	return result, err
}
