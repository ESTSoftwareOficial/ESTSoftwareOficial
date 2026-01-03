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
	createCourseWithRelations *application.CreateCourseWithRelations
}

func NewCreateCourseController(createCourseWithRelations *application.CreateCourseWithRelations) *CreateCourseController {
	return &CreateCourseController{createCourseWithRelations: createCourseWithRelations}
}

func (cc *CreateCourseController) Execute(c *gin.Context) {
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

	// Crea la entidad Course con IDs
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

	// Guarda y obtiene el curso con relaciones
	fmt.Println("Guardando curso en BD...")
	courseWithRelations, err := cc.createCourseWithRelations.Execute(course)
	if err != nil {
		fmt.Printf("ERROR al guardar curso: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// RESPUESTA con objetos relacionados (NUEVO)
	fmt.Println("Curso creado exitosamente")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Curso creado exitosamente",
		"course": dto.CourseResponse{
			ID:          courseWithRelations.ID,
			NameCourse:  courseWithRelations.NameCourse,
			Description: courseWithRelations.Description,
			Technology: dto.TechnologyDTO{
				TechnologyName:  courseWithRelations.TechnologyName,
				TechnologyImage: courseWithRelations.TechnologyImage,
			},
			Instructor: dto.InstructorDTO{
				InstructorName:  courseWithRelations.InstructorName,
				InstructorImage: courseWithRelations.InstructorImage,
			},
			CategoryName:  courseWithRelations.CategoryName,
			Level:         courseWithRelations.Level,
			ImageURL:      courseWithRelations.ImageURL,
			TotalModules:  courseWithRelations.TotalModules,
			AverageRating: courseWithRelations.AverageRating,
			TotalRatings:  courseWithRelations.TotalRatings,
			DurationHours: courseWithRelations.DurationHours,
			CreatedAt:     courseWithRelations.CreatedAt,
			UpdatedAt:     courseWithRelations.UpdatedAt,
			IsActive:      courseWithRelations.IsActive,
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