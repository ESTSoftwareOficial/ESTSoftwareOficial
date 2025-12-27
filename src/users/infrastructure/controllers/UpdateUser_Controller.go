package controllers

import (
	"estsoftwareoficial/src/core/cloudinary"
	"estsoftwareoficial/src/users/application"
	"estsoftwareoficial/src/users/domain/entities"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUser *application.UpdateUser
}

func NewUpdateUserController(updateUser *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{updateUser: updateUser}
}

func (uc *UpdateUserController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	firstName := c.PostForm("firstName")
	secondName := c.PostForm("secondName")
	lastName := c.PostForm("lastName")
	secondLastName := c.PostForm("secondLastName")
	email := c.PostForm("email")
	secondaryEmail := c.PostForm("secondaryEmail")
	roleIDStr := c.PostForm("roleId")

	if firstName == "" || lastName == "" || email == "" || roleIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obligatorios: firstName, lastName, email, roleId"})
		return
	}

	roleID, err := strconv.Atoi(roleIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roleId debe ser un número"})
		return
	}

	var secondaryEmailPtr *string
	if secondaryEmail != "" {
		secondaryEmailPtr = &secondaryEmail
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

		fmt.Println("Subiendo nueva foto a Cloudinary...")
		uploadedURL, err := cloudinary.UploadAvatar(fileBytes, header.Filename)
		if err != nil {
			fmt.Printf("ERROR al subir a Cloudinary: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al subir imagen"})
			return
		}

		fmt.Printf("Nueva foto subida: %s\n", uploadedURL)
		profilePhotoURL = &uploadedURL
	} else {
		fmt.Printf("No se recibió archivo profilePhoto (se mantendrá la foto actual)\n")
	}

	user := &entities.User{
		ID:             id,
		FirstName:      firstName,
		SecondName:     secondName,
		LastName:       lastName,
		SecondLastName: secondLastName,
		Email:          email,
		SecondaryEmail: secondaryEmailPtr,
		ProfilePhoto:   profilePhotoURL,
		RoleID:         roleID,
	}

	fmt.Println("Actualizando usuario en BD...")
	err = uc.updateUser.Execute(user)
	if err != nil {
		fmt.Printf("ERROR al actualizar usuario: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Usuario actualizado exitosamente")
	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario actualizado exitosamente",
	})
}
