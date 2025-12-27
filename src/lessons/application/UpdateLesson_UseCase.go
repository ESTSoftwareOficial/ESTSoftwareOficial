package application

import (
	"errors"

	"estsoftwareoficial/src/lessons/domain"
	"estsoftwareoficial/src/lessons/domain/entities"
)

type UpdateLesson struct {
	lessonRepo domain.LessonRepository
}

func NewUpdateLesson(lessonRepo domain.LessonRepository) *UpdateLesson {
	return &UpdateLesson{lessonRepo: lessonRepo}
}

func (ul *UpdateLesson) Execute(lesson *entities.Lesson) error {
	existingLesson, err := ul.lessonRepo.GetByID(lesson.ID)
	if err != nil {
		return err
	}
	if existingLesson == nil {
		return errors.New("lección no encontrada")
	}

	if lesson.Title == "" {
		return errors.New("el título es obligatorio")
	}

	if lesson.ContentType != "video" && lesson.ContentType != "markdown" &&
		lesson.ContentType != "quiz" && lesson.ContentType != "pdf" {
		return errors.New("tipo de contenido inválido, debe ser: video, markdown, quiz o pdf")
	}

	return ul.lessonRepo.Update(lesson)
}
