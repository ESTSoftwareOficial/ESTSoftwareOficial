package domain

import "estsoftwareoficial/src/comments/domain/entities"

type CommentRepository interface {
	Create(comment *entities.Comment) (*entities.Comment, error)
	GetByLessonID(lessonID, userID, limit, offset int) ([]*entities.Comment, error)
	GetByID(commentID int) (*entities.Comment, error)
	Update(comment *entities.Comment) error
	Delete(commentID int) error
	GetTotalByLesson(lessonID int) (int, error)
	AddLike(commentID, userID int) error
	RemoveLike(commentID, userID int) error
	GetLikesCount(commentID int) (int, error)
	UserHasLiked(commentID, userID int) (bool, error)
	GetUserInfo(userID int) (firstName, lastName string, profilePhoto *string, err error)
}
