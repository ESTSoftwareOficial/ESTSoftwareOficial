package dto

import "time"

type CourseRatingResponse struct {
    ID         int       `json:"id"`
    CourseName string    `json:"courseName"`  
    UserName   string    `json:"userName"`     
    Rating     int       `json:"rating"`
    Review     *string   `json:"review,omitempty"`
    CreatedAt  time.Time `json:"createdAt"`
    UpdatedAt  time.Time `json:"updatedAt"`
}