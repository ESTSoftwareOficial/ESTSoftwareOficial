package application

import (
	"estsoftwareoficial/src/course_ratings/domain"
	"estsoftwareoficial/src/course_ratings/domain/entities"
)

type GetCourseRatingById struct {
	courseRatingRepo domain.CourseRatingRepository
}

func NewGetCourseRatingById(courseRatingRepo domain.CourseRatingRepository) *GetCourseRatingById {
	return &GetCourseRatingById{courseRatingRepo: courseRatingRepo}
}

func (gcr *GetCourseRatingById) Execute(id int) (*entities.CourseRating, error) {
	return gcr.courseRatingRepo.GetByID(id)
}
