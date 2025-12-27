package dto

type ModuleRequest struct {
	CourseID    int     `json:"courseId" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description,omitempty"`
	OrderIndex  int     `json:"orderIndex" binding:"required"`
}

type UpdateModuleRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description,omitempty"`
	OrderIndex  int     `json:"orderIndex" binding:"required"`
}

type ReorderModulesRequest struct {
	Modules []ReorderModule `json:"modules" binding:"required"`
}

type ReorderModule struct {
	ID         int `json:"id" binding:"required"`
	OrderIndex int `json:"orderIndex" binding:"required"`
}
