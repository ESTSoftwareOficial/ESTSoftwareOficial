package application

import (
	"errors"

	"estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/courses/domain/entities"
)

type UpdateCourse struct {
	courseRepo domain.CourseRepository
}

func NewUpdateCourse(courseRepo domain.CourseRepository) *UpdateCourse {
	return &UpdateCourse{courseRepo: courseRepo}
}

func (uc *UpdateCourse) Execute(course *entities.Course) error {
	existingCourse, err := uc.courseRepo.GetByID(course.ID)
	if err != nil {
		return err
	}
	if existingCourse == nil {
		return errors.New("curso no encontrado")
	}

	if course.Level != "basico" && course.Level != "intermedio" && course.Level != "avanzado" {
		return errors.New("nivel inválido, debe ser: basico, intermedio o avanzado")
	}

	// Si no se envía nueva imagen, mantener la anterior
	if course.ImageURL == nil || *course.ImageURL == "" {
		course.ImageURL = existingCourse.ImageURL
	}

	return uc.courseRepo.Update(course)
}
