package entities

import "time"

// CourseWithRelations incluye datos de las tablas relacionadas
type CourseWithRelations struct {
	// Datos del curso
	ID            int
	NameCourse    string
	Description   string
	Level         string
	ImageURL      *string
	TotalModules  int
	AverageRating float64
	TotalRatings  int
	DurationHours *float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	IsActive      bool

	// Datos relacionados
	TechnologyName  string
	TechnologyImage string
	InstructorName  string
	InstructorImage *string
	CategoryName    string
}