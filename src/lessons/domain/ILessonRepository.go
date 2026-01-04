package domain

import "estsoftwareoficial/src/lessons/domain/entities"

type LessonRepository interface {
	Save(lesson *entities.Lesson) (*entities.Lesson, error)
	GetByID(id int) (*entities.Lesson, error)
	GetDetailByID(id int) (*entities.LessonDetail, error)
	GetAll() ([]*entities.Lesson, error)
	GetByModule(moduleID int) ([]*entities.Lesson, error)
	Update(lesson *entities.Lesson) error
	Delete(id int) error
	UpdateOrderIndex(id int, orderIndex int) error
}
