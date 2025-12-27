package dto

type LessonResourceRequest struct {
	LessonID       int     `json:"lessonId" binding:"required"`
	ResourceTypeID int     `json:"resourceTypeId" binding:"required"`
	URL            string  `json:"url" binding:"required"`
	Title          *string `json:"title,omitempty"`
}

type UpdateLessonResourceRequest struct {
	ResourceTypeID int     `json:"resourceTypeId" binding:"required"`
	URL            string  `json:"url" binding:"required"`
	Title          *string `json:"title,omitempty"`
}
