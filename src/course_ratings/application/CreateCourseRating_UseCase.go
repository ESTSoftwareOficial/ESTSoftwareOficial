package application

import (
	"errors"

	"estsoftwareoficial/src/course_ratings/domain"
	"estsoftwareoficial/src/course_ratings/domain/entities"
	coursesDomain "estsoftwareoficial/src/courses/domain"
)

type CreateCourseRating struct {
	courseRatingRepo domain.CourseRatingRepository
	courseRepo       coursesDomain.CourseRepository
}

func NewCreateCourseRating(courseRatingRepo domain.CourseRatingRepository, courseRepo coursesDomain.CourseRepository) *CreateCourseRating {
	return &CreateCourseRating{
		courseRatingRepo: courseRatingRepo,
		courseRepo:       courseRepo,
	}
}

func (ccr *CreateCourseRating) Execute(courseRating *entities.CourseRating) (*entities.CourseRating, error) {
	if courseRating.Rating < 1 || courseRating.Rating > 5 {
		return nil, errors.New("la calificación debe estar entre 1 y 5")
	}

	if courseRating.CourseID == 0 {
		return nil, errors.New("el ID del curso es obligatorio")
	}

	if courseRating.UserID == 0 {
		return nil, errors.New("el ID del usuario es obligatorio")
	}

	existingRating, _ := ccr.courseRatingRepo.GetByUserAndCourse(courseRating.UserID, courseRating.CourseID)
	if existingRating != nil {
		return nil, errors.New("ya has calificado este curso, usa actualizar en su lugar")
	}

	savedRating, err := ccr.courseRatingRepo.Save(courseRating)
	if err != nil {
		return nil, err
	}

	average, total, err := ccr.courseRatingRepo.CalculateAverageRating(courseRating.CourseID)
	if err != nil {
		return nil, errors.New("error al calcular el promedio de calificaciones")
	}

	err = ccr.courseRepo.UpdateRating(courseRating.CourseID, average, total)
	if err != nil {
		return nil, errors.New("error al actualizar el rating del curso")
	}

	return savedRating, nil
}
