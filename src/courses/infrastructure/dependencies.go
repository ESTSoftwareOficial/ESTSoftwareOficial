package infrastructure

import (
	"estsoftwareoficial/src/core"
	"estsoftwareoficial/src/courses/application"
	"estsoftwareoficial/src/courses/infrastructure/adapters"
	"estsoftwareoficial/src/courses/infrastructure/controllers"
)

type DependenciesCourses struct {
	CreateCourseController           *controllers.CreateCourseController
	GetAllCoursesController          *controllers.GetAllCoursesController
	GetCourseByIdController          *controllers.GetCourseByIdController
	GetCoursesByInstructorController *controllers.GetCoursesByInstructorController
	GetCoursesByCategoryController   *controllers.GetCoursesByCategoryController
	GetCoursesByTechnologyController *controllers.GetCoursesByTechnologyController
	UpdateCourseController           *controllers.UpdateCourseController
	DeleteCourseController           *controllers.DeleteCourseController
	SearchCoursesController          *controllers.SearchCoursesController
}

func InitCourses() *DependenciesCourses {
	conn := core.GetDBPool()
	courseRepo := adapters.NewPostgreSQL(conn.DB)

	createCourse := application.NewCreateCourse(courseRepo)
	getAllCourses := application.NewGetAllCourses(courseRepo)
	getCourseById := application.NewGetCourseById(courseRepo)
	getCoursesByInstructor := application.NewGetCoursesByInstructor(courseRepo)
	getCoursesByCategory := application.NewGetCoursesByCategory(courseRepo)
	getCoursesByTechnology := application.NewGetCoursesByTechnology(courseRepo)
	updateCourse := application.NewUpdateCourse(courseRepo)
	deleteCourse := application.NewDeleteCourse(courseRepo)
	searchCourses := application.NewSearchCourses(courseRepo)

	return &DependenciesCourses{
		CreateCourseController:           controllers.NewCreateCourseController(createCourse),
		GetAllCoursesController:          controllers.NewGetAllCoursesController(getAllCourses),
		GetCourseByIdController:          controllers.NewGetCourseByIdController(getCourseById),
		GetCoursesByInstructorController: controllers.NewGetCoursesByInstructorController(getCoursesByInstructor),
		GetCoursesByCategoryController:   controllers.NewGetCoursesByCategoryController(getCoursesByCategory),
		GetCoursesByTechnologyController: controllers.NewGetCoursesByTechnologyController(getCoursesByTechnology),
		UpdateCourseController:           controllers.NewUpdateCourseController(updateCourse),
		DeleteCourseController:           controllers.NewDeleteCourseController(deleteCourse),
		SearchCoursesController:          controllers.NewSearchCoursesController(searchCourses),
	}
}
