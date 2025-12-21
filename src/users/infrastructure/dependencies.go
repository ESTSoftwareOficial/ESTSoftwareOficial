package infrastructure

import (
	"estsoftwareoficial/src/core"
	"estsoftwareoficial/src/users/application"
	"estsoftwareoficial/src/users/infrastructure/adapters"
	"estsoftwareoficial/src/users/infrastructure/controllers"
)

type DependenciesUsers struct {
	AuthController          *controllers.AuthController
	CreateUserController    *controllers.CreateUserController
	GetAllUsersController   *controllers.GetAllUsersController
	GetUserByIdController   *controllers.GetUserByIdController
	GetTotalUsersController *controllers.GetTotalUsersController
	UpdateUserController    *controllers.UpdateUserController
	DeleteUserController    *controllers.DeleteUserController
	OAuthController         *controllers.OAuthController
}

func InitUsers() *DependenciesUsers {
	conn := core.GetDBPool()
	userRepo := adapters.NewPostgreSQL(conn.DB)

	authService := application.NewAuthService(userRepo)
	getAllUsers := application.NewGetAllUsers(userRepo)
	getUserById := application.NewGetUserById(userRepo)
	getTotalUsers := application.NewGetTotalUsers(userRepo)
	updateUser := application.NewUpdateUser(userRepo)
	deleteUser := application.NewDeleteUser(userRepo)
	oauthService := application.NewOAuthService(userRepo)

	return &DependenciesUsers{
		AuthController:          controllers.NewAuthController(authService),
		CreateUserController:    controllers.NewCreateUserController(authService),
		GetAllUsersController:   controllers.NewGetAllUsersController(getAllUsers),
		GetUserByIdController:   controllers.NewGetUserByIdController(getUserById),
		GetTotalUsersController: controllers.NewGetTotalUsersController(getTotalUsers),
		UpdateUserController:    controllers.NewUpdateUserController(updateUser),
		DeleteUserController:    controllers.NewDeleteUserController(deleteUser),
		OAuthController:         controllers.NewOAuthController(oauthService),
	}
}
