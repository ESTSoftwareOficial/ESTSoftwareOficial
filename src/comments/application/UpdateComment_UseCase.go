package application

import (
	"errors"
	"estsoftwareoficial/src/comments/domain"
)

type UpdateComment struct {
	commentRepo domain.CommentRepository
}

func NewUpdateComment(commentRepo domain.CommentRepository) *UpdateComment {
	return &UpdateComment{commentRepo: commentRepo}
}

func (uc *UpdateComment) Execute(commentID, userID int, newComment string) error {
	if newComment == "" {
		return errors.New("el comentario no puede estar vacío")
	}

	if len(newComment) > 1000 {
		return errors.New("el comentario no puede exceder 1000 caracteres")
	}

	comment, err := uc.commentRepo.GetByID(commentID)
	if err != nil {
		return err
	}

	if comment == nil {
		return errors.New("comentario no encontrado")
	}

	if comment.UserID != userID {
		return errors.New("no tienes permiso para editar este comentario")
	}

	if comment.IsDeleted {
		return errors.New("no se puede editar un comentario eliminado")
	}

	comment.Comment = newComment
	return uc.commentRepo.Update(comment)
}
