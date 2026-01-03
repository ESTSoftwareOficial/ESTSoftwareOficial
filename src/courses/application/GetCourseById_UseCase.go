package application

import (
	"estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/courses/domain/entities"
)

type GetCourseById struct {
	courseRepo domain.CourseRepository
}

func NewGetCourseById(courseRepo domain.CourseRepository) *GetCourseById {
	return &GetCourseById{courseRepo: courseRepo}
}

func (gc *GetCourseById) Execute(id int) (*entities.CourseWithRelations, error) {
	return gc.courseRepo.GetByIDWithRelations(id)
}