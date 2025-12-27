package application

import (
	"estsoftwareoficial/src/course_ratings/domain"
	"estsoftwareoficial/src/course_ratings/domain/entities"
)

type GetAllCourseRatings struct {
	courseRatingRepo domain.CourseRatingRepository
}

func NewGetAllCourseRatings(courseRatingRepo domain.CourseRatingRepository) *GetAllCourseRatings {
	return &GetAllCourseRatings{courseRatingRepo: courseRatingRepo}
}

func (gcr *GetAllCourseRatings) Execute() ([]*entities.CourseRating, error) {
	return gcr.courseRatingRepo.GetAll()
}
