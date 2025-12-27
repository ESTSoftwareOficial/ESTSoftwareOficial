package dto

import "time"

type CourseResponse struct {
	ID            int       `json:"id"`
	NameCourse    string    `json:"nameCourse"`
	Description   string    `json:"description"`
	TechnologyID  int       `json:"technologyId"`
	InstructorID  int       `json:"instructorId"`
	CategoryID    int       `json:"categoryId"`
	Level         string    `json:"level"`
	ImageURL      *string   `json:"imageUrl,omitempty"`
	TotalModules  int       `json:"totalModules"`
	AverageRating float64   `json:"averageRating"`
	TotalRatings  int       `json:"totalRatings"`
	DurationHours *float64  `json:"durationHours,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	IsActive      bool      `json:"isActive"`
}
