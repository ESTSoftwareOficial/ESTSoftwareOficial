package application

import (
	"estsoftwareoficial/src/lessons/domain"
	"estsoftwareoficial/src/lessons/domain/entities"
)

type GetLessonDetail struct {
	lessonRepo domain.LessonRepository
}

func NewGetLessonDetail(lessonRepo domain.LessonRepository) *GetLessonDetail {
	return &GetLessonDetail{lessonRepo: lessonRepo}
}

func (gl *GetLessonDetail) Execute(id int) (*entities.LessonDetail, error) {
	return gl.lessonRepo.GetDetailByID(id)
}