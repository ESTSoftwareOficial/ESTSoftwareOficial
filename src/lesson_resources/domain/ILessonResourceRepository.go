package domain

import "estsoftwareoficial/src/lesson_resources/domain/entities"

type LessonResourceRepository interface {
	Save(lessonResource *entities.LessonResource) (*entities.LessonResource, error)
	GetByID(id int) (*entities.LessonResource, error)
	GetAll() ([]*entities.LessonResource, error)
	GetByLesson(lessonID int) ([]*entities.LessonResource, error)
	Update(lessonResource *entities.LessonResource) error
	Delete(id int) error
}
