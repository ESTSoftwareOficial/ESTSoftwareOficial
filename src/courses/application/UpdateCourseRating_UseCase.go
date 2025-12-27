package application

import (
	"estsoftwareoficial/src/courses/domain"
)

type UpdateCourseRating struct {
	courseRepo domain.CourseRepository
}

func NewUpdateCourseRating(courseRepo domain.CourseRepository) *UpdateCourseRating {
	return &UpdateCourseRating{courseRepo: courseRepo}
}

func (ucr *UpdateCourseRating) Execute(courseID int, averageRating float64, totalRatings int) error {
	return ucr.courseRepo.UpdateRating(courseID, averageRating, totalRatings)
}
