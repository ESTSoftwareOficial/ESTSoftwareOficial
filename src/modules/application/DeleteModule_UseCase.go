package application

import (
	"errors"

	coursesDomain "estsoftwareoficial/src/courses/domain"
	"estsoftwareoficial/src/modules/domain"
)

type DeleteModule struct {
	moduleRepo domain.ModuleRepository
	courseRepo coursesDomain.CourseRepository
}

func NewDeleteModule(moduleRepo domain.ModuleRepository, courseRepo coursesDomain.CourseRepository) *DeleteModule {
	return &DeleteModule{
		moduleRepo: moduleRepo,
		courseRepo: courseRepo,
	}
}

func (dm *DeleteModule) Execute(id int) error {
	existingModule, err := dm.moduleRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingModule == nil {
		return errors.New("módulo no encontrado")
	}

	courseID := existingModule.CourseID

	err = dm.moduleRepo.Delete(id)
	if err != nil {
		return err
	}

	totalModules, err := dm.moduleRepo.GetTotalModulesByCourse(courseID)
	if err != nil {
		return err
	}

	err = dm.courseRepo.UpdateTotalModules(courseID, totalModules)
	if err != nil {
		return errors.New("error al actualizar total de módulos del curso")
	}

	return nil
}
