package application

import (
	"estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/courses/domain/entities"
)

type GetAllCourses struct {
	courseRepo domain.CourseRepository
}

func NewGetAllCourses(courseRepo domain.CourseRepository) *GetAllCourses {
	return &GetAllCourses{courseRepo: courseRepo}
}

func (gc *GetAllCourses) Execute() ([]*entities.Course, error) {
	return gc.courseRepo.GetAll()
}
