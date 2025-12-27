package dto

type CourseRequest struct {
	NameCourse    string   `json:"nameCourse" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	TechnologyID  int      `json:"technologyId" binding:"required"`
	InstructorID  int      `json:"instructorId" binding:"required"`
	CategoryID    int      `json:"categoryId" binding:"required"`
	Level         string   `json:"level" binding:"required,oneof=basico intermedio avanzado"`
	ImageURL      *string  `json:"imageUrl,omitempty"`
	DurationHours *float64 `json:"durationHours,omitempty"`
}

type UpdateCourseRequest struct {
	NameCourse    string   `json:"nameCourse" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	TechnologyID  int      `json:"technologyId" binding:"required"`
	CategoryID    int      `json:"categoryId" binding:"required"`
	Level         string   `json:"level" binding:"required,oneof=basico intermedio avanzado"`
	ImageURL      *string  `json:"imageUrl,omitempty"`
	DurationHours *float64 `json:"durationHours,omitempty"`
	IsActive      *bool    `json:"isActive,omitempty"`
}
