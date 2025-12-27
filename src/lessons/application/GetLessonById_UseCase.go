package application

import (
	"estsoftwareoficial/src/lessons/domain"
	"estsoftwareoficial/src/lessons/domain/entities"
)

type GetLessonById struct {
	lessonRepo domain.LessonRepository
}

func NewGetLessonById(lessonRepo domain.LessonRepository) *GetLessonById {
	return &GetLessonById{lessonRepo: lessonRepo}
}

func (gl *GetLessonById) Execute(id int) (*entities.Lesson, error) {
	return gl.lessonRepo.GetByID(id)
}
