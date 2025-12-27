package infrastructure

import (
	"estsoftwareoficial/src/core"
	"estsoftwareoficial/src/lesson_resources/application"
	"estsoftwareoficial/src/lesson_resources/infrastructure/adapters"
	"estsoftwareoficial/src/lesson_resources/infrastructure/controllers"
)

type DependenciesLessonResources struct {
	CreateLessonResourceController       *controllers.CreateLessonResourceController
	GetAllLessonResourcesController      *controllers.GetAllLessonResourcesController
	GetLessonResourceByIdController      *controllers.GetLessonResourceByIdController
	GetLessonResourcesByLessonController *controllers.GetLessonResourcesByLessonController
	UpdateLessonResourceController       *controllers.UpdateLessonResourceController
	DeleteLessonResourceController       *controllers.DeleteLessonResourceController
}

func InitLessonResources() *DependenciesLessonResources {
	conn := core.GetDBPool()
	lessonResourceRepo := adapters.NewPostgreSQL(conn.DB)

	createLessonResource := application.NewCreateLessonResource(lessonResourceRepo)
	getAllLessonResources := application.NewGetAllLessonResources(lessonResourceRepo)
	getLessonResourceById := application.NewGetLessonResourceById(lessonResourceRepo)
	getLessonResourcesByLesson := application.NewGetLessonResourcesByLesson(lessonResourceRepo)
	updateLessonResource := application.NewUpdateLessonResource(lessonResourceRepo)
	deleteLessonResource := application.NewDeleteLessonResource(lessonResourceRepo)

	return &DependenciesLessonResources{
		CreateLessonResourceController:       controllers.NewCreateLessonResourceController(createLessonResource),
		GetAllLessonResourcesController:      controllers.NewGetAllLessonResourcesController(getAllLessonResources),
		GetLessonResourceByIdController:      controllers.NewGetLessonResourceByIdController(getLessonResourceById),
		GetLessonResourcesByLessonController: controllers.NewGetLessonResourcesByLessonController(getLessonResourcesByLesson),
		UpdateLessonResourceController:       controllers.NewUpdateLessonResourceController(updateLessonResource),
		DeleteLessonResourceController:       controllers.NewDeleteLessonResourceController(deleteLessonResource),
	}
}
