package application

import (
	"errors"

	"estsoftwareoficial/src/lessons/domain"
)

type ReorderLessons struct {
	lessonRepo domain.LessonRepository
}

func NewReorderLessons(lessonRepo domain.LessonRepository) *ReorderLessons {
	return &ReorderLessons{lessonRepo: lessonRepo}
}

type LessonOrder struct {
	ID         int
	OrderIndex int
}

func (rl *ReorderLessons) Execute(lessons []LessonOrder) error {
	if len(lessons) == 0 {
		return errors.New("no se proporcionaron lecciones para reordenar")
	}

	for _, lesson := range lessons {
		err := rl.lessonRepo.UpdateOrderIndex(lesson.ID, lesson.OrderIndex)
		if err != nil {
			return err
		}
	}

	return nil
}
