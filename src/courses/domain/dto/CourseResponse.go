package dto

import "time"

// DTOs para las entidades relacionadas
type TechnologyDTO struct {
	TechnologyName  string `json:"technologyName"`
	TechnologyImage string `json:"technologyImage"`
}

type InstructorDTO struct {
	InstructorName  string  `json:"instructorName"`
	InstructorImage *string `json:"instructorImage,omitempty"`
}

// CourseResponse actualizado con relaciones
type CourseResponse struct {
	ID            int            `json:"id"`
	NameCourse    string         `json:"nameCourse"`
	Description   string         `json:"description"`
	Technology    TechnologyDTO  `json:"technology"`
	Instructor    InstructorDTO  `json:"instructor"`
	CategoryName  string         `json:"categoryName"`
	Level         string         `json:"level"`
	ImageURL      *string        `json:"imageUrl,omitempty"`
	TotalModules  int            `json:"totalModules"`
	AverageRating float64        `json:"averageRating"`
	TotalRatings  int            `json:"totalRatings"`
	DurationHours *float64       `json:"durationHours,omitempty"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	IsActive      bool           `json:"isActive"`
}