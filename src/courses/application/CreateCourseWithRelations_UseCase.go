package application

import (
	"errors"
	"time"

	"estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/courses/domain/entities"
)

type CreateCourseWithRelations struct {
	courseRepo domain.CourseRepository
}

func NewCreateCourseWithRelations(courseRepo domain.CourseRepository) *CreateCourseWithRelations {
	return &CreateCourseWithRelations{courseRepo: courseRepo}
}

func (cc *CreateCourseWithRelations) Execute(course *entities.Course) (*entities.CourseWithRelations, error) {
	// Validaciones básicas
	if course.NameCourse == "" {
		return nil, errors.New("el nombre del curso es obligatorio")
	}

	if course.Description == "" {
		return nil, errors.New("la descripción es obligatoria")
	}

	if course.Level != "basico" && course.Level != "intermedio" && course.Level != "avanzado" {
		return nil, errors.New("nivel inválido, debe ser: basico, intermedio o avanzado")
	}

	// Establecer valores por defecto
	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()
	course.IsActive = true
	course.TotalModules = 0
	course.AverageRating = 0.00
	course.TotalRatings = 0

	// Guardar el curso normalmente (con IDs)
	savedCourse, err := cc.courseRepo.Save(course)
	if err != nil {
		return nil, err
	}

	// Obtener el curso con las relaciones (hacer JOIN)
	return cc.courseRepo.GetByIDWithRelations(savedCourse.ID)
}