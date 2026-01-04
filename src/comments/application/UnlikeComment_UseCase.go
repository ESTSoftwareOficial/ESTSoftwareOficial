package application

import "estsoftwareoficial/src/comments/domain"

type UnlikeComment struct {
	commentRepo domain.CommentRepository
}

func NewUnlikeComment(commentRepo domain.CommentRepository) *UnlikeComment {
	return &UnlikeComment{commentRepo: commentRepo}
}

func (uc *UnlikeComment) Execute(commentID, userID int) (int, error) {
	err := uc.commentRepo.RemoveLike(commentID, userID)
	if err != nil {
		return 0, err
	}

	likesCount, err := uc.commentRepo.GetLikesCount(commentID)
	if err != nil {
		return 0, err
	}

	return likesCount, nil
}
