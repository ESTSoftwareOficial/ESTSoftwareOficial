package application

import (
	"errors"
	"estsoftwareoficial/src/comments/domain"
)

type DeleteComment struct {
	commentRepo domain.CommentRepository
}

func NewDeleteComment(commentRepo domain.CommentRepository) *DeleteComment {
	return &DeleteComment{commentRepo: commentRepo}
}

func (dc *DeleteComment) Execute(commentID, userID int) error {
	comment, err := dc.commentRepo.GetByID(commentID)
	if err != nil {
		return err
	}

	if comment == nil {
		return errors.New("comentario no encontrado")
	}

	if comment.UserID != userID {
		return errors.New("no tienes permiso para eliminar este comentario")
	}

	return dc.commentRepo.Delete(commentID)
}
