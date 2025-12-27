package controllers

import (
	"estsoftwareoficial/src/core/cloudinary"
	"estsoftwareoficial/src/users/application"
	"estsoftwareoficial/src/users/domain/dto"
	"estsoftwareoficial/src/users/domain/entities"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	authService *application.AuthService
}

func NewCreateUserController(authService *application.AuthService) *CreateUserController {
	return &CreateUserController{authService: authService}
}

func (uc *CreateUserController) Execute(c *gin.Context) {
	firstName := c.PostForm("firstName")
	secondName := c.PostForm("secondName")
	lastName := c.PostForm("lastName")
	secondLastName := c.PostForm("secondLastName")
	email := c.PostForm("email")
	password := c.PostForm("password")
	roleIDStr := c.PostForm("roleId")

	fmt.Printf("Datos recibidos - Email: %s, RoleID: %s\n", email, roleIDStr)

	if firstName == "" || lastName == "" || email == "" || password == "" || roleIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obligatorios: firstName, lastName, email, password, roleId"})
		return
	}

	roleID, err := strconv.Atoi(roleIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roleId debe ser un número"})
		return
	}

	var profilePhotoURL *string
	fmt.Println("Intentando leer archivo profilePhoto...")
	file, header, err := c.Request.FormFile("profilePhoto")
	if err == nil {
		fmt.Printf("Archivo recibido: %s, Tamaño: %d bytes\n", header.Filename, header.Size)
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("ERROR al leer bytes del archivo: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar imagen"})
			return
		}

		fmt.Printf("Bytes leídos: %d\n", len(fileBytes))
		fmt.Println("Subiendo a Cloudinary...")

		uploadedURL, err := cloudinary.UploadAvatar(fileBytes, header.Filename)
		if err != nil {
			fmt.Printf("ERROR al subir a Cloudinary: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al subir imagen: %v", err)})
			return
		}

		fmt.Printf("Imagen subida exitosamente: %s\n", uploadedURL)
		profilePhotoURL = &uploadedURL
	} else {
		fmt.Printf("No se recibió archivo profilePhoto: %v\n", err)
	}

	user := &entities.User{
		FirstName:      firstName,
		SecondName:     secondName,
		LastName:       lastName,
		SecondLastName: secondLastName,
		Email:          email,
		Password:       &password,
		ProfilePhoto:   profilePhotoURL,
		RoleID:         roleID,
	}

	fmt.Println("Registrando usuario en BD...")
	savedUser, err := uc.authService.Register(user)
	if err != nil {
		fmt.Printf("ERROR al registrar usuario: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Usuario creado exitosamente")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado exitosamente",
		"user": dto.UserResponse{
			ID:               savedUser.ID,
			FirstName:        savedUser.FirstName,
			SecondName:       savedUser.SecondName,
			LastName:         savedUser.LastName,
			SecondLastName:   savedUser.SecondLastName,
			Email:            savedUser.Email,
			SecondaryEmail:   savedUser.SecondaryEmail,
			ProfilePhoto:     savedUser.ProfilePhoto,
			RegistrationDate: savedUser.RegistrationDate,
			RoleID:           savedUser.RoleID,
		},
	})
}
