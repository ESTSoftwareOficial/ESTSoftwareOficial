package application

import "estsoftwareoficial/src/comments/domain"

type LikeComment struct {
	commentRepo domain.CommentRepository
}

func NewLikeComment(commentRepo domain.CommentRepository) *LikeComment {
	return &LikeComment{commentRepo: commentRepo}
}

func (lc *LikeComment) Execute(commentID, userID int) (int, error) {
	err := lc.commentRepo.AddLike(commentID, userID)
	if err != nil {
		return 0, err
	}

	likesCount, err := lc.commentRepo.GetLikesCount(commentID)
	if err != nil {
		return 0, err
	}

	return likesCount, nil
}
