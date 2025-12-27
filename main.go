package main

import (
	categoryInfrastructure "estsoftwareoficial/src/categories/infrastructure"
	categoryRoutes "estsoftwareoficial/src/categories/infrastructure/routes"
	"estsoftwareoficial/src/core/cloudinary"
	courseRatingInfrastructure "estsoftwareoficial/src/course_ratings/infrastructure"
	courseRatingRoutes "estsoftwareoficial/src/course_ratings/infrastructure/routes"
	courseInfrastructure "estsoftwareoficial/src/courses/infrastructure"
	courseRoutes "estsoftwareoficial/src/courses/infrastructure/routes"
	lessonResourceInfrastructure "estsoftwareoficial/src/lesson_resources/infrastructure"
	lessonResourceRoutes "estsoftwareoficial/src/lesson_resources/infrastructure/routes"
	lessonInfrastructure "estsoftwareoficial/src/lessons/infrastructure"
	lessonRoutes "estsoftwareoficial/src/lessons/infrastructure/routes"
	moduleInfrastructure "estsoftwareoficial/src/modules/infrastructure"
	moduleRoutes "estsoftwareoficial/src/modules/infrastructure/routes"
	resourceTypeInfrastructure "estsoftwareoficial/src/resource_types/infrastructure"
	resourceTypeRoutes "estsoftwareoficial/src/resource_types/infrastructure/routes"
	technologyInfrastructure "estsoftwareoficial/src/technologies/infrastructure"
	technologyRoutes "estsoftwareoficial/src/technologies/infrastructure/routes"
	userInfrastructure "estsoftwareoficial/src/users/infrastructure"
	userRoutes "estsoftwareoficial/src/users/infrastructure/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cloudinary.InitCloudinary()
	userDeps := userInfrastructure.InitUsers()
	categoryDeps := categoryInfrastructure.InitCategories()
	technologyDeps := technologyInfrastructure.InitTechnologies()
	resourceTypeDeps := resourceTypeInfrastructure.InitResourceTypes()
	courseDeps := courseInfrastructure.InitCourses()
	moduleDeps := moduleInfrastructure.InitModules()
	lessonDeps := lessonInfrastructure.InitLessons()
	lessonResourceDeps := lessonResourceInfrastructure.InitLessonResources()
	courseRatingDeps := courseRatingInfrastructure.InitCourseRatings()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	userRoutes.ConfigureUserRoutes(
		router,
		userDeps.AuthController,
		userDeps.CreateUserController,
		userDeps.GetAllUsersController,
		userDeps.GetUserByIdController,
		userDeps.GetTotalUsersController,
		userDeps.UpdateUserController,
		userDeps.DeleteUserController,
		userDeps.OAuthController,
	)

	categoryRoutes.ConfigureCategoryRoutes(
		router,
		categoryDeps.CreateCategoryController,
		categoryDeps.GetAllCategoriesController,
		categoryDeps.GetCategoryByIdController,
		categoryDeps.UpdateCategoryController,
		categoryDeps.DeleteCategoryController,
	)

	technologyRoutes.ConfigureTechnologyRoutes(
		router,
		technologyDeps.CreateTechnologyController,
		technologyDeps.GetAllTechnologiesController,
		technologyDeps.GetTechnologyByIdController,
		technologyDeps.UpdateTechnologyController,
		technologyDeps.DeleteTechnologyController,
	)

	resourceTypeRoutes.ConfigureResourceTypeRoutes(
		router,
		resourceTypeDeps.CreateResourceTypeController,
		resourceTypeDeps.GetAllResourceTypesController,
		resourceTypeDeps.GetResourceTypeByIdController,
		resourceTypeDeps.UpdateResourceTypeController,
		resourceTypeDeps.DeleteResourceTypeController,
	)

	courseRoutes.ConfigureCourseRoutes(
		router,
		courseDeps.CreateCourseController,
		courseDeps.GetAllCoursesController,
		courseDeps.GetCourseByIdController,
		courseDeps.GetCoursesByInstructorController,
		courseDeps.GetCoursesByCategoryController,
		courseDeps.GetCoursesByTechnologyController,
		courseDeps.UpdateCourseController,
		courseDeps.DeleteCourseController,
		courseDeps.SearchCoursesController,
	)

	moduleRoutes.ConfigureModuleRoutes(
		router,
		moduleDeps.CreateModuleController,
		moduleDeps.GetAllModulesController,
		moduleDeps.GetModuleByIdController,
		moduleDeps.GetModulesByCourseController,
		moduleDeps.UpdateModuleController,
		moduleDeps.DeleteModuleController,
		moduleDeps.ReorderModulesController,
	)

	lessonRoutes.ConfigureLessonRoutes(
		router,
		lessonDeps.CreateLessonController,
		lessonDeps.GetAllLessonsController,
		lessonDeps.GetLessonByIdController,
		lessonDeps.GetLessonsByModuleController,
		lessonDeps.UpdateLessonController,
		lessonDeps.DeleteLessonController,
		lessonDeps.ReorderLessonsController,
	)

	lessonResourceRoutes.ConfigureLessonResourceRoutes(
		router,
		lessonResourceDeps.CreateLessonResourceController,
		lessonResourceDeps.GetAllLessonResourcesController,
		lessonResourceDeps.GetLessonResourceByIdController,
		lessonResourceDeps.GetLessonResourcesByLessonController,
		lessonResourceDeps.UpdateLessonResourceController,
		lessonResourceDeps.DeleteLessonResourceController,
	)

	courseRatingRoutes.ConfigureCourseRatingRoutes(
		router,
		courseRatingDeps.CreateCourseRatingController,
		courseRatingDeps.GetAllCourseRatingsController,
		courseRatingDeps.GetCourseRatingByIdController,
		courseRatingDeps.GetCourseRatingsByCourseController,
		courseRatingDeps.GetCourseRatingByUserAndCourseController,
		courseRatingDeps.UpdateCourseRatingController,
		courseRatingDeps.DeleteCourseRatingController,
	)

	log.Println("Servidor corriendo en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
