package application

import (
	"estsoftwareoficial/src/lessons/domain"
	"estsoftwareoficial/src/lessons/domain/entities"
)

type GetLessonsByModule struct {
	lessonRepo domain.LessonRepository
}

func NewGetLessonsByModule(lessonRepo domain.LessonRepository) *GetLessonsByModule {
	return &GetLessonsByModule{lessonRepo: lessonRepo}
}

func (gl *GetLessonsByModule) Execute(moduleID int) ([]*entities.Lesson, error) {
	return gl.lessonRepo.GetByModule(moduleID)
}
