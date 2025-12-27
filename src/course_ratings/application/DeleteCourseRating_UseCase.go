package application

import (
	"errors"

	"estsoftwareoficial/src/course_ratings/domain"
	coursesDomain "estsoftwareoficial/src/courses/domain"
)

type DeleteCourseRating struct {
	courseRatingRepo domain.CourseRatingRepository
	courseRepo       coursesDomain.CourseRepository
}

func NewDeleteCourseRating(courseRatingRepo domain.CourseRatingRepository, courseRepo coursesDomain.CourseRepository) *DeleteCourseRating {
	return &DeleteCourseRating{
		courseRatingRepo: courseRatingRepo,
		courseRepo:       courseRepo,
	}
}

func (dcr *DeleteCourseRating) Execute(id int) error {
	existingRating, err := dcr.courseRatingRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingRating == nil {
		return errors.New("calificación no encontrada")
	}

	courseID := existingRating.CourseID

	err = dcr.courseRatingRepo.Delete(id)
	if err != nil {
		return err
	}

	average, total, err := dcr.courseRatingRepo.CalculateAverageRating(courseID)
	if err != nil {
		return errors.New("error al calcular el promedio de calificaciones")
	}

	err = dcr.courseRepo.UpdateRating(courseID, average, total)
	if err != nil {
		return errors.New("error al actualizar el rating del curso")
	}

	return nil
}
