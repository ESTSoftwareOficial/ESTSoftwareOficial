package entities

import "time"

type Comment struct {
	ID        int       `json:"id"`
	LessonID  int       `json:"lessonId"`
	UserID    int       `json:"userId"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsEdited  bool      `json:"isEdited"`
	IsDeleted bool      `json:"isDeleted"`
}
