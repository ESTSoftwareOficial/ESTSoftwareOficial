package dto

type ResourceTypeRequest struct {
	Name    string  `json:"name" binding:"required"`
	IconURL *string `json:"iconUrl,omitempty"`
}
