package application

import (
	"estsoftwareoficial/src/comments/domain"
	"estsoftwareoficial/src/comments/domain/dto"
	"math"
)

type GetLessonComments struct {
	commentRepo domain.CommentRepository
}

func NewGetLessonComments(commentRepo domain.CommentRepository) *GetLessonComments {
	return &GetLessonComments{commentRepo: commentRepo}
}

func (glc *GetLessonComments) Execute(lessonID, userID, page, limit int) (*dto.CommentListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	offset := (page - 1) * limit

	comments, err := glc.commentRepo.GetByLessonID(lessonID, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := glc.commentRepo.GetTotalByLesson(lessonID)
	if err != nil {
		return nil, err
	}

	var commentResponses []dto.CommentResponse
	for _, comment := range comments {
		firstName, lastName, profilePhoto, err := glc.commentRepo.GetUserInfo(comment.UserID)
		if err != nil {
			continue
		}

		likesCount, _ := glc.commentRepo.GetLikesCount(comment.ID)
		userHasLiked, _ := glc.commentRepo.UserHasLiked(comment.ID, userID)

		commentResponses = append(commentResponses, dto.CommentResponse{
			ID:       comment.ID,
			LessonID: comment.LessonID,
			User: dto.UserInfo{
				ID:           comment.UserID,
				FirstName:    firstName,
				LastName:     lastName,
				ProfilePhoto: profilePhoto,
			},
			Comment:      comment.Comment,
			CreatedAt:    comment.CreatedAt,
			UpdatedAt:    comment.UpdatedAt,
			IsEdited:     comment.IsEdited,
			LikesCount:   likesCount,
			UserHasLiked: userHasLiked,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &dto.CommentListResponse{
		Comments: commentResponses,
		Pagination: dto.PaginationInfo{
			CurrentPage:   page,
			TotalPages:    totalPages,
			TotalComments: total,
			Limit:         limit,
		},
	}, nil
}
