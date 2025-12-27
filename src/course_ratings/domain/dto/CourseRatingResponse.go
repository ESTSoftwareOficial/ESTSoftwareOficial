package dto

import "time"

type CourseRatingResponse struct {
	ID        int       `json:"id"`
	CourseID  int       `json:"courseId"`
	UserID    int       `json:"userId"`
	Rating    int       `json:"rating"`
	Review    *string   `json:"review,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
