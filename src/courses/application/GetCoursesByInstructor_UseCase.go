package application

import (
	"estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/courses/domain/entities"
)

type GetCoursesByInstructor struct {
	courseRepo domain.CourseRepository
}

func NewGetCoursesByInstructor(courseRepo domain.CourseRepository) *GetCoursesByInstructor {
	return &GetCoursesByInstructor{courseRepo: courseRepo}
}

func (gc *GetCoursesByInstructor) Execute(instructorID int) ([]*entities.Course, error) {
	return gc.courseRepo.GetByInstructor(instructorID)
}
