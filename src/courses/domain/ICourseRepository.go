package domain

import "estsoftwareoficial/src/courses/domain/entities"

type CourseRepository interface {
	Save(course *entities.Course) (*entities.Course, error)
	GetByID(id int) (*entities.Course, error)
	GetAll() ([]*entities.Course, error)
	GetByInstructor(instructorID int) ([]*entities.Course, error)
	GetByCategory(categoryID int) ([]*entities.Course, error)
	GetByTechnology(technologyID int) ([]*entities.Course, error)
	Update(course *entities.Course) error
	Delete(id int) error
	UpdateTotalModules(courseID int, totalModules int) error
	UpdateRating(courseID int, averageRating float64, totalRatings int) error
	Search(keyword string, categoryID *int, technologyID *int, level *string, minRating *float64) ([]*entities.Course, error)
	
	// Nuevos métodos con JOINs
	GetByIDWithRelations(id int) (*entities.CourseWithRelations, error)
	GetAllWithRelations() ([]*entities.CourseWithRelations, error)
	GetByInstructorWithRelations(instructorID int) ([]*entities.CourseWithRelations, error)
	GetByCategoryWithRelations(categoryID int) ([]*entities.CourseWithRelations, error)
	GetByTechnologyWithRelations(technologyID int) ([]*entities.CourseWithRelations, error)
	SearchWithRelations(keyword string, categoryID *int, technologyID *int, level *string, minRating *float64) ([]*entities.CourseWithRelations, error)
}