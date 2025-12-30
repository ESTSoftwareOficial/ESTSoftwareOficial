package entities

import (
	"os"
	"time"
)

type Lesson struct {
	ID              int       `json:"id"`
	ModuleID        int       `json:"moduleId"`
	Title           string    `json:"title"`
	ContentType     string    `json:"contentType"`
	BunnyLibraryID  *string   `json:"bunnyLibraryId,omitempty"`
	BunnyVideoID    *string   `json:"bunnyVideoId,omitempty"`
	BodyText        *string   `json:"bodyText,omitempty"`
	DurationMinutes int       `json:"durationMinutes"`
	OrderIndex      int       `json:"orderIndex"`
	IsPreview       bool      `json:"isPreview"`
	CreatedAt       time.Time `json:"createdAt"`
}

func (l *Lesson) GetVideoURL() *string {
	if l.BunnyLibraryID != nil && l.BunnyVideoID != nil {
		pullZoneURL := os.Getenv("BUNNY_PULL_ZONE_URL")
		url := pullZoneURL + "/embed/" + *l.BunnyLibraryID + "/" + *l.BunnyVideoID
		return &url
	}
	return nil
}
