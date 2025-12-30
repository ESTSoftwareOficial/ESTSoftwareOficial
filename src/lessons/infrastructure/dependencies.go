package infrastructure

import (
	"estsoftwareoficial/src/core"
	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/infrastructure/adapters"
	"estsoftwareoficial/src/lessons/infrastructure/controllers"
)

type DependenciesLessons struct {
	CreateLessonController       *controllers.CreateLessonController
	GetAllLessonsController      *controllers.GetAllLessonsController
	GetLessonByIdController      *controllers.GetLessonByIdController
	GetLessonsByModuleController *controllers.GetLessonsByModuleController
	UpdateLessonController       *controllers.UpdateLessonController
	DeleteLessonController       *controllers.DeleteLessonController
	ReorderLessonsController     *controllers.ReorderLessonsController
}

func InitLessons() *DependenciesLessons {
	conn := core.GetDBPool()
	lessonRepo := adapters.NewPostgreSQL(conn.DB)

	createLesson := application.NewCreateLesson(lessonRepo)
	getAllLessons := application.NewGetAllLessons(lessonRepo)
	getLessonById := application.NewGetLessonById(lessonRepo)
	getLessonsByModule := application.NewGetLessonsByModule(lessonRepo)
	updateLesson := application.NewUpdateLesson(lessonRepo)
	deleteLesson := application.NewDeleteLesson(lessonRepo)
	reorderLessons := application.NewReorderLessons(lessonRepo)

	return &DependenciesLessons{
		CreateLessonController:       controllers.NewCreateLessonController(createLesson),
		GetAllLessonsController:      controllers.NewGetAllLessonsController(getAllLessons),
		GetLessonByIdController:      controllers.NewGetLessonByIdController(getLessonById),
		GetLessonsByModuleController: controllers.NewGetLessonsByModuleController(getLessonsByModule),
		UpdateLessonController:       controllers.NewUpdateLessonController(updateLesson, getLessonById),
		DeleteLessonController:       controllers.NewDeleteLessonController(deleteLesson, getLessonById),
		ReorderLessonsController:     controllers.NewReorderLessonsController(reorderLessons),
	}
}
