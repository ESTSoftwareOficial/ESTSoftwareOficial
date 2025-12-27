package application

import (
	"errors"

	"estsoftwareoficial/src/lesson_resources/domain"
	"estsoftwareoficial/src/lesson_resources/domain/entities"
)

type CreateLessonResource struct {
	lessonResourceRepo domain.LessonResourceRepository
}

func NewCreateLessonResource(lessonResourceRepo domain.LessonResourceRepository) *CreateLessonResource {
	return &CreateLessonResource{lessonResourceRepo: lessonResourceRepo}
}

func (clr *CreateLessonResource) Execute(lessonResource *entities.LessonResource) (*entities.LessonResource, error) {
	if lessonResource.URL == "" {
		return nil, errors.New("la URL es obligatoria")
	}

	if lessonResource.LessonID == 0 {
		return nil, errors.New("el ID de la lección es obligatorio")
	}

	if lessonResource.ResourceTypeID == 0 {
		return nil, errors.New("el ID del tipo de recurso es obligatorio")
	}

	return clr.lessonResourceRepo.Save(lessonResource)
}
