package dto

type CreateCommentRequest struct {
	Comment string `json:"comment" binding:"required,min=1,max=1000"`
}

type UpdateCommentRequest struct {
	Comment string `json:"comment" binding:"required,min=1,max=1000"`
}
