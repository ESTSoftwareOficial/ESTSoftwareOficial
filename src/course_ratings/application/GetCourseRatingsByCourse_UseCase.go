package application

import (
	"estsoftwareoficial/src/course_ratings/domain"
	"estsoftwareoficial/src/course_ratings/domain/entities"
)

type GetCourseRatingsByCourse struct {
	courseRatingRepo domain.CourseRatingRepository
}

func NewGetCourseRatingsByCourse(courseRatingRepo domain.CourseRatingRepository) *GetCourseRatingsByCourse {
	return &GetCourseRatingsByCourse{courseRatingRepo: courseRatingRepo}
}

func (gcr *GetCourseRatingsByCourse) Execute(courseID int) ([]*entities.CourseRating, error) {
	return gcr.courseRatingRepo.GetByCourse(courseID)
}
