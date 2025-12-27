package application

import (
	"estsoftwareoficial/src/lessons/domain"
	"estsoftwareoficial/src/lessons/domain/entities"
)

type GetAllLessons struct {
	lessonRepo domain.LessonRepository
}

func NewGetAllLessons(lessonRepo domain.LessonRepository) *GetAllLessons {
	return &GetAllLessons{lessonRepo: lessonRepo}
}

func (gl *GetAllLessons) Execute() ([]*entities.Lesson, error) {
	return gl.lessonRepo.GetAll()
}
