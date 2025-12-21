package controllers

import (
	"estsoftwareoficial/src/core/security"
	"estsoftwareoficial/src/users/application"
	"estsoftwareoficial/src/users/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *application.AuthService
}

func NewAuthController(authService *application.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	user, err := ac.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := security.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	refreshToken, err := security.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar refresh token"})
		return
	}

	security.SetAuthCookie(c.Writer, accessToken)
	security.SetRefreshCookie(c.Writer, refreshToken)

	c.JSON(http.StatusOK, dto.LoginResponse{
		Message: "Login exitoso",
		User: dto.UserResponse{
			ID:               user.ID,
			FirstName:        user.FirstName,
			SecondName:       user.SecondName,
			LastName:         user.LastName,
			SecondLastName:   user.SecondLastName,
			Email:            user.Email,
			SecondaryEmail:   user.SecondaryEmail,
			RegistrationDate: user.RegistrationDate,
			RoleID:           user.RoleID,
			OAuthProvider:    user.OAuthProvider,
		},
	})
}

func (ac *AuthController) Logout(c *gin.Context) {
	security.ClearAuthCookies(c.Writer)
	c.JSON(http.StatusOK, gin.H{"message": "Logout exitoso"})
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token no encontrado"})
		return
	}

	claims, err := security.ValidateRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	user, err := ac.authService.GetUserByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	newAccessToken, err := security.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	security.SetAuthCookie(c.Writer, newAccessToken)
	c.JSON(http.StatusOK, gin.H{"message": "Token renovado"})
}

func (ac *AuthController) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autenticado"})
		return
	}

	user, err := ac.authService.GetUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": dto.UserResponse{
			ID:               user.ID,
			FirstName:        user.FirstName,
			SecondName:       user.SecondName,
			LastName:         user.LastName,
			SecondLastName:   user.SecondLastName,
			Email:            user.Email,
			SecondaryEmail:   user.SecondaryEmail,
			RegistrationDate: user.RegistrationDate,
			RoleID:           user.RoleID,
			OAuthProvider:    user.OAuthProvider,
		},
	})
}

func (ac *AuthController) VerifyToken(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	roleID, _ := c.Get("role_id")

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"user": gin.H{
			"id":     userID,
			"email":  email,
			"roleId": roleID,
		},
	})
}
