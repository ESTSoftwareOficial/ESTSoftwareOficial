package application

import (
	"errors"

	coursesDomain "estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/modules/domain"
	"estsoftwareoficial/src/modules/domain/entities"
)

type CreateModule struct {
	moduleRepo domain.ModuleRepository
	courseRepo coursesDomain.CourseRepository
}

func NewCreateModule(moduleRepo domain.ModuleRepository, courseRepo coursesDomain.CourseRepository) *CreateModule {
	return &CreateModule{
		moduleRepo: moduleRepo,
		courseRepo: courseRepo,
	}
}

func (cm *CreateModule) Execute(module *entities.Module) (*entities.Module, error) {
	if module.Title == "" {
		return nil, errors.New("el título es obligatorio")
	}

	if module.CourseID == 0 {
		return nil, errors.New("el ID del curso es obligatorio")
	}

	course, err := cm.courseRepo.GetByID(module.CourseID)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, errors.New("el curso no existe")
	}

	savedModule, err := cm.moduleRepo.Save(module)
	if err != nil {
		return nil, err
	}

	totalModules, err := cm.moduleRepo.GetTotalModulesByCourse(module.CourseID)
	if err != nil {
		return nil, err
	}

	err = cm.courseRepo.UpdateTotalModules(module.CourseID, totalModules)
	if err != nil {
		return nil, errors.New("error al actualizar total de módulos del curso")
	}

	return savedModule, nil
}
