package dto

type LessonRequest struct {
	ModuleID        int     `json:"moduleId" binding:"required"`
	Title           string  `json:"title" binding:"required"`
	ContentType     string  `json:"contentType" binding:"required,oneof=video markdown quiz pdf"`
	ContentURL      *string `json:"contentUrl,omitempty"`
	BodyText        *string `json:"bodyText,omitempty"`
	DurationMinutes int     `json:"durationMinutes"`
	OrderIndex      int     `json:"orderIndex" binding:"required"`
	IsPreview       bool    `json:"isPreview"`
}

type UpdateLessonRequest struct {
	Title           string  `json:"title" binding:"required"`
	ContentType     string  `json:"contentType" binding:"required,oneof=video markdown quiz pdf"`
	ContentURL      *string `json:"contentUrl,omitempty"`
	BodyText        *string `json:"bodyText,omitempty"`
	DurationMinutes int     `json:"durationMinutes"`
	OrderIndex      int     `json:"orderIndex" binding:"required"`
	IsPreview       bool    `json:"isPreview"`
}

type ReorderLessonsRequest struct {
	Lessons []ReorderLesson `json:"lessons" binding:"required"`
}

type ReorderLesson struct {
	ID         int `json:"id" binding:"required"`
	OrderIndex int `json:"orderIndex" binding:"required"`
}
