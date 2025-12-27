package application

import (
	"estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/courses/domain/entities"
)

type GetCoursesByTechnology struct {
	courseRepo domain.CourseRepository
}

func NewGetCoursesByTechnology(courseRepo domain.CourseRepository) *GetCoursesByTechnology {
	return &GetCoursesByTechnology{courseRepo: courseRepo}
}

func (gc *GetCoursesByTechnology) Execute(technologyID int) ([]*entities.Course, error) {
	return gc.courseRepo.GetByTechnology(technologyID)
}
