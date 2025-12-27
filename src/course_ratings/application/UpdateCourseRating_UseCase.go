package application

import (
	"errors"

	"estsoftwareoficial/src/course_ratings/domain"
	"estsoftwareoficial/src/course_ratings/domain/entities"
	coursesDomain "estsoftwareoficial/src/courses/domain"
)

type UpdateCourseRating struct {
	courseRatingRepo domain.CourseRatingRepository
	courseRepo       coursesDomain.CourseRepository
}

func NewUpdateCourseRating(courseRatingRepo domain.CourseRatingRepository, courseRepo coursesDomain.CourseRepository) *UpdateCourseRating {
	return &UpdateCourseRating{
		courseRatingRepo: courseRatingRepo,
		courseRepo:       courseRepo,
	}
}

func (ucr *UpdateCourseRating) Execute(courseRating *entities.CourseRating) error {
	existingRating, err := ucr.courseRatingRepo.GetByID(courseRating.ID)
	if err != nil {
		return err
	}
	if existingRating == nil {
		return errors.New("calificación no encontrada")
	}

	if courseRating.Rating < 1 || courseRating.Rating > 5 {
		return errors.New("la calificación debe estar entre 1 y 5")
	}

	err = ucr.courseRatingRepo.Update(courseRating)
	if err != nil {
		return err
	}

	average, total, err := ucr.courseRatingRepo.CalculateAverageRating(existingRating.CourseID)
	if err != nil {
		return errors.New("error al calcular el promedio de calificaciones")
	}

	err = ucr.courseRepo.UpdateRating(existingRating.CourseID, average, total)
	if err != nil {
		return errors.New("error al actualizar el rating del curso")
	}

	return nil
}
