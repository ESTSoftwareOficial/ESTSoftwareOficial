package application

import (
	"errors"
	"time"

	"estsoftwareoficial/src/lessons/domain"
	"estsoftwareoficial/src/lessons/domain/entities"
)

type CreateLesson struct {
	lessonRepo domain.LessonRepository
}

func NewCreateLesson(lessonRepo domain.LessonRepository) *CreateLesson {
	return &CreateLesson{lessonRepo: lessonRepo}
}

func (cl *CreateLesson) Execute(lesson *entities.Lesson) (*entities.Lesson, error) {
	if lesson.Title == "" {
		return nil, errors.New("el título es obligatorio")
	}

	if lesson.ModuleID == 0 {
		return nil, errors.New("el ID del módulo es obligatorio")
	}

	if lesson.ContentType != "video" && lesson.ContentType != "markdown" &&
		lesson.ContentType != "quiz" && lesson.ContentType != "pdf" {
		return nil, errors.New("tipo de contenido inválido, debe ser: video, markdown, quiz o pdf")
	}

	lesson.CreatedAt = time.Now()

	return cl.lessonRepo.Save(lesson)
}
