package dto

type ResourceTypeResponse struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	IconURL *string `json:"iconUrl,omitempty"`
}
