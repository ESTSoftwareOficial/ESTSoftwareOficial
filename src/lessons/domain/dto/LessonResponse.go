package dto

import "time"

type LessonResponse struct {
	ID              int       `json:"id"`
	ModuleID        int       `json:"moduleId"`
	Title           string    `json:"title"`
	ContentType     string    `json:"contentType"`
	ContentURL      *string   `json:"contentUrl,omitempty"`
	BodyText        *string   `json:"bodyText,omitempty"`
	DurationMinutes int       `json:"durationMinutes"`
	OrderIndex      int       `json:"orderIndex"`
	IsPreview       bool      `json:"isPreview"`
	CreatedAt       time.Time `json:"createdAt"`
}
