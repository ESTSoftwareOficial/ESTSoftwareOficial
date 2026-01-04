package infrastructure

import (
	"estsoftwareoficial/src/comments/application"
	"estsoftwareoficial/src/comments/infrastructure/adapters"
	"estsoftwareoficial/src/comments/infrastructure/controllers"
	"estsoftwareoficial/src/core"
)

type DependenciesComments struct {
	CreateCommentController     *controllers.CreateCommentController
	GetLessonCommentsController *controllers.GetLessonCommentsController
	UpdateCommentController     *controllers.UpdateCommentController
	DeleteCommentController     *controllers.DeleteCommentController
	LikeCommentController       *controllers.LikeCommentController
	UnlikeCommentController     *controllers.UnlikeCommentController
}

func InitComments() *DependenciesComments {
	conn := core.GetDBPool()
	commentRepo := adapters.NewPostgreSQL(conn.DB)

	createComment := application.NewCreateComment(commentRepo)
	getLessonComments := application.NewGetLessonComments(commentRepo)
	updateComment := application.NewUpdateComment(commentRepo)
	deleteComment := application.NewDeleteComment(commentRepo)
	likeComment := application.NewLikeComment(commentRepo)
	unlikeComment := application.NewUnlikeComment(commentRepo)

	return &DependenciesComments{
		CreateCommentController:     controllers.NewCreateCommentController(createComment),
		GetLessonCommentsController: controllers.NewGetLessonCommentsController(getLessonComments),
		UpdateCommentController:     controllers.NewUpdateCommentController(updateComment),
		DeleteCommentController:     controllers.NewDeleteCommentController(deleteComment),
		LikeCommentController:       controllers.NewLikeCommentController(likeComment),
		UnlikeCommentController:     controllers.NewUnlikeCommentController(unlikeComment),
	}
}
