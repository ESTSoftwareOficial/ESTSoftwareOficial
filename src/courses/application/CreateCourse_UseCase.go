package application

import (
	"errors"
	"time"

	"estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/courses/domain/entities"
)

type CreateCourse struct {
	courseRepo domain.CourseRepository
}

func NewCreateCourse(courseRepo domain.CourseRepository) *CreateCourse {
	return &CreateCourse{courseRepo: courseRepo}
}

func (cc *CreateCourse) Execute(course *entities.Course) (*entities.Course, error) {
	if course.NameCourse == "" {
		return nil, errors.New("el nombre del curso es obligatorio")
	}

	if course.Description == "" {
		return nil, errors.New("la descripción es obligatoria")
	}

	if course.Level != "basico" && course.Level != "intermedio" && course.Level != "avanzado" {
		return nil, errors.New("nivel inválido, debe ser: basico, intermedio o avanzado")
	}

	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()
	course.IsActive = true
	course.TotalModules = 0
	course.AverageRating = 0.00
	course.TotalRatings = 0

	return cc.courseRepo.Save(course)
}
