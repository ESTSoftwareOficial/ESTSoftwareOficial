package infrastructure

import (
	"estsoftwareoficial/src/core"
	"estsoftwareoficial/src/course_ratings/application"
	"estsoftwareoficial/src/course_ratings/infrastructure/adapters"
	"estsoftwareoficial/src/course_ratings/infrastructure/controllers"
	coursesDomain "estsoftwareoficial/src/courses/domain"
	coursesAdapters "estsoftwareoficial/src/courses/infrastructure/adapters"
)

type DependenciesCourseRatings struct {
	CreateCourseRatingController             *controllers.CreateCourseRatingController
	GetAllCourseRatingsController            *controllers.GetAllCourseRatingsController
	GetCourseRatingByIdController            *controllers.GetCourseRatingByIdController
	GetCourseRatingsByCourseController       *controllers.GetCourseRatingsByCourseController
	GetCourseRatingByUserAndCourseController *controllers.GetCourseRatingByUserAndCourseController
	UpdateCourseRatingController             *controllers.UpdateCourseRatingController
	DeleteCourseRatingController             *controllers.DeleteCourseRatingController
}

func InitCourseRatings() *DependenciesCourseRatings {
	conn := core.GetDBPool()
	courseRatingRepo := adapters.NewPostgreSQL(conn.DB)

	var courseRepo coursesDomain.CourseRepository = coursesAdapters.NewPostgreSQL(conn.DB)

	createCourseRating := application.NewCreateCourseRating(courseRatingRepo, courseRepo)
	getAllCourseRatings := application.NewGetAllCourseRatings(courseRatingRepo)
	getCourseRatingById := application.NewGetCourseRatingById(courseRatingRepo)
	getCourseRatingsByCourse := application.NewGetCourseRatingsByCourse(courseRatingRepo)
	getCourseRatingByUserAndCourse := application.NewGetCourseRatingByUserAndCourse(courseRatingRepo)
	updateCourseRating := application.NewUpdateCourseRating(courseRatingRepo, courseRepo)
	deleteCourseRating := application.NewDeleteCourseRating(courseRatingRepo, courseRepo)

	return &DependenciesCourseRatings{
		CreateCourseRatingController:             controllers.NewCreateCourseRatingController(createCourseRating),
		GetAllCourseRatingsController:            controllers.NewGetAllCourseRatingsController(getAllCourseRatings),
		GetCourseRatingByIdController:            controllers.NewGetCourseRatingByIdController(getCourseRatingById),
		GetCourseRatingsByCourseController:       controllers.NewGetCourseRatingsByCourseController(getCourseRatingsByCourse),
		GetCourseRatingByUserAndCourseController: controllers.NewGetCourseRatingByUserAndCourseController(getCourseRatingByUserAndCourse),
		UpdateCourseRatingController:             controllers.NewUpdateCourseRatingController(updateCourseRating),
		DeleteCourseRatingController:             controllers.NewDeleteCourseRatingController(deleteCourseRating),
	}
}
