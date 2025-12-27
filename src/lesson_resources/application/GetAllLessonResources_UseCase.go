package application

import (
	"estsoftwareoficial/src/lesson_resources/domain"
	"estsoftwareoficial/src/lesson_resources/domain/entities"
)

type GetAllLessonResources struct {
	lessonResourceRepo domain.LessonResourceRepository
}

func NewGetAllLessonResources(lessonResourceRepo domain.LessonResourceRepository) *GetAllLessonResources {
	return &GetAllLessonResources{lessonResourceRepo: lessonResourceRepo}
}

func (glr *GetAllLessonResources) Execute() ([]*entities.LessonResource, error) {
	return glr.lessonResourceRepo.GetAll()
}
