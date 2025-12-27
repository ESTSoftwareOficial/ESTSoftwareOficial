package application

import (
	"errors"

	"estsoftwareoficial/src/lesson_resources/domain"
)

type DeleteLessonResource struct {
	lessonResourceRepo domain.LessonResourceRepository
}

func NewDeleteLessonResource(lessonResourceRepo domain.LessonResourceRepository) *DeleteLessonResource {
	return &DeleteLessonResource{lessonResourceRepo: lessonResourceRepo}
}

func (dlr *DeleteLessonResource) Execute(id int) error {
	existingLessonResource, err := dlr.lessonResourceRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingLessonResource == nil {
		return errors.New("recurso de lección no encontrado")
	}

	return dlr.lessonResourceRepo.Delete(id)
}
