package dto

type ModuleResponse struct {
	ID          int     `json:"id"`
	CourseID    int     `json:"courseId"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	OrderIndex  int     `json:"orderIndex"`
}
