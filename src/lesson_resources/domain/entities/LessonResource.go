package entities

type LessonResource struct {
	ID             int     `json:"id"`
	LessonID       int     `json:"lessonId"`
	ResourceTypeID int     `json:"resourceTypeId"`
	URL            string  `json:"url"`
	Title          *string `json:"title,omitempty"`
}
