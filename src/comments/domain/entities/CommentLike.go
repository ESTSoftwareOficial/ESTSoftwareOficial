package entities

import "time"

type CommentLike struct {
	ID        int       `json:"id"`
	CommentID int       `json:"commentId"`
	UserID    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
