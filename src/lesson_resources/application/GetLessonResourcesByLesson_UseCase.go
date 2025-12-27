package application

import (
	"estsoftwareoficial/src/lesson_resources/domain"
	"estsoftwareoficial/src/lesson_resources/domain/entities"
)

type GetLessonResourcesByLesson struct {
	lessonResourceRepo domain.LessonResourceRepository
}

func NewGetLessonResourcesByLesson(lessonResourceRepo domain.LessonResourceRepository) *GetLessonResourcesByLesson {
	return &GetLessonResourcesByLesson{lessonResourceRepo: lessonResourceRepo}
}

func (glr *GetLessonResourcesByLesson) Execute(lessonID int) ([]*entities.LessonResource, error) {
	return glr.lessonResourceRepo.GetByLesson(lessonID)
}
