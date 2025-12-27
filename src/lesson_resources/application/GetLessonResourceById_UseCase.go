package application

import (
	"estsoftwareoficial/src/lesson_resources/domain"
	"estsoftwareoficial/src/lesson_resources/domain/entities"
)

type GetLessonResourceById struct {
	lessonResourceRepo domain.LessonResourceRepository
}

func NewGetLessonResourceById(lessonResourceRepo domain.LessonResourceRepository) *GetLessonResourceById {
	return &GetLessonResourceById{lessonResourceRepo: lessonResourceRepo}
}

func (glr *GetLessonResourceById) Execute(id int) (*entities.LessonResource, error) {
	return glr.lessonResourceRepo.GetByID(id)
}
