package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"estsoftwareoficial/src/core/bunny"
	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateLessonController struct {
	updateLesson  *application.UpdateLesson
	getLessonById *application.GetLessonById
}

func NewUpdateLessonController(updateLesson *application.UpdateLesson, getLessonById *application.GetLessonById) *UpdateLessonController {
	return &UpdateLessonController{
		updateLesson:  updateLesson,
		getLessonById: getLessonById,
	}
}

func (ul *UpdateLessonController) Execute(c *gin.Context) {
	fmt.Println("=== INICIO UPDATE LESSON ===")

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	existingLesson, err := ul.getLessonById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lección no encontrada"})
		return
	}

	title := c.PostForm("title")
	contentType := c.PostForm("contentType")
	bodyText := c.PostForm("bodyText")
	durationMinutes := c.PostForm("durationMinutes")
	orderIndex := c.PostForm("orderIndex")
	isPreview := c.PostForm("isPreview")

	if title == "" || contentType == "" || orderIndex == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obligatorios: title, contentType, orderIndex"})
		return
	}

	durationMinutesInt := 0
	if durationMinutes != "" {
		durationMinutesInt, _ = parseToInt(durationMinutes)
	}

	orderIndexInt, err := parseToInt(orderIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "orderIndex debe ser un número"})
		return
	}

	isPreviewBool := isPreview == "true"

	var bodyTextPtr *string
	if bodyText != "" {
		bodyTextPtr = &bodyText
	}

	lesson := &entities.Lesson{
		ID:              id,
		Title:           title,
		ContentType:     contentType,
		BodyText:        bodyTextPtr,
		DurationMinutes: durationMinutesInt,
		OrderIndex:      orderIndexInt,
		IsPreview:       isPreviewBool,
		BunnyLibraryID:  existingLesson.BunnyLibraryID,
		BunnyVideoID:    existingLesson.BunnyVideoID,
	}

	if contentType == "video" {
		file, header, err := c.Request.FormFile("video")
		if err == nil {
			defer file.Close()
			fmt.Printf("Nuevo video recibido: %s, Tamaño: %d bytes\n", header.Filename, header.Size)

			if existingLesson.BunnyVideoID != nil {
				fmt.Printf("Eliminando video anterior de Bunny.net: %s\n", *existingLesson.BunnyVideoID)
				_ = bunny.DeleteVideo(*existingLesson.BunnyVideoID)
			}

			fileBytes, err := io.ReadAll(file)
			if err != nil {
				fmt.Printf("ERROR al leer bytes del video: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar video"})
				return
			}

			fmt.Println("Subiendo nuevo video a Bunny.net...")
			videoID, err := bunny.UploadVideo(fileBytes, header.Filename, title)
			if err != nil {
				fmt.Printf("ERROR al subir a Bunny.net: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al subir video"})
				return
			}

			fmt.Printf("Nuevo video subido exitosamente. Video ID: %s\n", videoID)

			libraryID := os.Getenv("BUNNY_LIBRARY_ID")
			lesson.BunnyLibraryID = &libraryID
			lesson.BunnyVideoID = &videoID
		}
	}

	err = ul.updateLesson.Execute(lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lección actualizada exitosamente",
	})
}
