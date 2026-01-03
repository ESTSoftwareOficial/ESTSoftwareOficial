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

// Retorna cursos con relaciones
func (gc *GetAllCourses) Execute() ([]*entities.CourseWithRelations, error) {
	return gc.courseRepo.GetAllWithRelations()
}