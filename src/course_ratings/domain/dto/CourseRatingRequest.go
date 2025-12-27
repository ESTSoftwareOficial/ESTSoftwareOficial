package dto

type CourseRatingRequest struct {
	CourseID int     `json:"courseId" binding:"required"`
	UserID   int     `json:"userId" binding:"required"`
	Rating   int     `json:"rating" binding:"required,min=1,max=5"`
	Review   *string `json:"review,omitempty"`
}

type UpdateCourseRatingRequest struct {
	Rating int     `json:"rating" binding:"required,min=1,max=5"`
	Review *string `json:"review,omitempty"`
}
