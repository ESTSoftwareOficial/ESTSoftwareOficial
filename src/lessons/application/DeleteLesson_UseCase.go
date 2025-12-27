package application

import (
	"errors"

	"estsoftwareoficial/src/lessons/domain"
)

type DeleteLesson struct {
	lessonRepo domain.LessonRepository
}

func NewDeleteLesson(lessonRepo domain.LessonRepository) *DeleteLesson {
	return &DeleteLesson{lessonRepo: lessonRepo}
}

func (dl *DeleteLesson) Execute(id int) error {
	existingLesson, err := dl.lessonRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingLesson == nil {
		return errors.New("lección no encontrada")
	}

	return dl.lessonRepo.Delete(id)
}
