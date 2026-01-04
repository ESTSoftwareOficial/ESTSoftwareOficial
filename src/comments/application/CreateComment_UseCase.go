package application

import (
	"errors"
	"estsoftwareoficial/src/comments/domain"
	"estsoftwareoficial/src/comments/domain/entities"
)

type CreateComment struct {
	commentRepo domain.CommentRepository
}

func NewCreateComment(commentRepo domain.CommentRepository) *CreateComment {
	return &CreateComment{commentRepo: commentRepo}
}

func (cc *CreateComment) Execute(lessonID, userID int, commentText string) (*entities.Comment, error) {
	if commentText == "" {
		return nil, errors.New("el comentario no puede estar vacío")
	}

	if len(commentText) > 1000 {
		return nil, errors.New("el comentario no puede exceder 1000 caracteres")
	}

	comment := &entities.Comment{
		LessonID: lessonID,
		UserID:   userID,
		Comment:  commentText,
	}

	return cc.commentRepo.Create(comment)
}
