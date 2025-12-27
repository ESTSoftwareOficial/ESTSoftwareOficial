package application

import (
	"estsoftwareoficial/src/course_ratings/domain"
	"estsoftwareoficial/src/course_ratings/domain/entities"
)

type GetCourseRatingByUserAndCourse struct {
	courseRatingRepo domain.CourseRatingRepository
}

func NewGetCourseRatingByUserAndCourse(courseRatingRepo domain.CourseRatingRepository) *GetCourseRatingByUserAndCourse {
	return &GetCourseRatingByUserAndCourse{courseRatingRepo: courseRatingRepo}
}

func (gcr *GetCourseRatingByUserAndCourse) Execute(userID int, courseID int) (*entities.CourseRating, error) {
	return gcr.courseRatingRepo.GetByUserAndCourse(userID, courseID)
}
