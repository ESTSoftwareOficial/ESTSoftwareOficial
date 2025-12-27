package application

import (
	"estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/courses/domain/entities"
)

type SearchCourses struct {
	courseRepo domain.CourseRepository
}

func NewSearchCourses(courseRepo domain.CourseRepository) *SearchCourses {
	return &SearchCourses{courseRepo: courseRepo}
}

func (sc *SearchCourses) Execute(keyword string, categoryID *int, technologyID *int, level *string, minRating *float64) ([]*entities.Course, error) {
	return sc.courseRepo.Search(keyword, categoryID, technologyID, level, minRating)
}
