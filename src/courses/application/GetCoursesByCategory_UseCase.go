package application

import (
	"estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/courses/domain/entities"
)

type GetCoursesByCategory struct {
	courseRepo domain.CourseRepository
}

func NewGetCoursesByCategory(courseRepo domain.CourseRepository) *GetCoursesByCategory {
	return &GetCoursesByCategory{courseRepo: courseRepo}
}

func (gc *GetCoursesByCategory) Execute(categoryID int) ([]*entities.CourseWithRelations, error) {
	return gc.courseRepo.GetByCategoryWithRelations(categoryID)
}