package main

import (
	"estsoftwareoficial/src/core/cloudinary"
	"estsoftwareoficial/src/users/infrastructure"
	"estsoftwareoficial/src/users/infrastructure/routes"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("API REST - ESTSoftware")

	cloudinary.InitCloudinary()

	userDeps := infrastructure.InitUsers()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	routes.ConfigureUserRoutes(
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

	log.Println("Servidor corriendo en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
