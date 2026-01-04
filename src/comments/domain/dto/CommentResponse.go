package dto

import "time"

type CommentResponse struct {
	ID           int       `json:"id"`
	LessonID     int       `json:"lessonId"`
	User         UserInfo  `json:"user"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	IsEdited     bool      `json:"isEdited"`
	LikesCount   int       `json:"likesCount"`
	UserHasLiked bool      `json:"userHasLiked"`
}

type UserInfo struct {
	ID           int     `json:"id"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	ProfilePhoto *string `json:"profilePhoto,omitempty"`
}

type CommentListResponse struct {
	Comments   []CommentResponse `json:"comments"`
	Pagination PaginationInfo    `json:"pagination"`
}

type PaginationInfo struct {
	CurrentPage   int `json:"currentPage"`
	TotalPages    int `json:"totalPages"`
	TotalComments int `json:"totalComments"`
	Limit         int `json:"limit"`
}

type LikeResponse struct {
	CommentID  int  `json:"commentId"`
	Liked      bool `json:"liked"`
	LikesCount int  `json:"likesCount"`
}
