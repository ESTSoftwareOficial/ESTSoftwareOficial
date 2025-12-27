package application

import (
	"errors"

	"estsoftwareoficial/src/core/cloudinary"
	"estsoftwareoficial/src/courses/domain"
)

type DeleteCourse struct {
	courseRepo domain.CourseRepository
}

func NewDeleteCourse(courseRepo domain.CourseRepository) *DeleteCourse {
	return &DeleteCourse{courseRepo: courseRepo}
}

func (dc *DeleteCourse) Execute(id int) error {
	existingCourse, err := dc.courseRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingCourse == nil {
		return errors.New("curso no encontrado")
	}

	// Eliminar imagen de Cloudinary si existe
	if existingCourse.ImageURL != nil && *existingCourse.ImageURL != "" {
		if err := cloudinary.DeleteImage(*existingCourse.ImageURL); err != nil {
			return errors.New("error al eliminar imagen del curso de Cloudinary")
		}
	}

	return dc.courseRepo.Delete(id)
}
