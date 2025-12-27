package application

import (
	"errors"

	"estsoftwareoficial/src/lesson_resources/domain"
	"estsoftwareoficial/src/lesson_resources/domain/entities"
)

type UpdateLessonResource struct {
	lessonResourceRepo domain.LessonResourceRepository
}

func NewUpdateLessonResource(lessonResourceRepo domain.LessonResourceRepository) *UpdateLessonResource {
	return &UpdateLessonResource{lessonResourceRepo: lessonResourceRepo}
}

func (ulr *UpdateLessonResource) Execute(lessonResource *entities.LessonResource) error {
	existingLessonResource, err := ulr.lessonResourceRepo.GetByID(lessonResource.ID)
	if err != nil {
		return err
	}
	if existingLessonResource == nil {
		return errors.New("recurso de lección no encontrado")
	}

	if lessonResource.URL == "" {
		return errors.New("la URL es obligatoria")
	}

	if lessonResource.ResourceTypeID == 0 {
		return errors.New("el ID del tipo de recurso es obligatorio")
	}

	return ulr.lessonResourceRepo.Update(lessonResource)
}
