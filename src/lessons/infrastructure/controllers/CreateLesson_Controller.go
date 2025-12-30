package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"estsoftwareoficial/src/core/bunny"
	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/domain/dto"
	"estsoftwareoficial/src/lessons/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateLessonController struct {
	createLesson *application.CreateLesson
}

func NewCreateLessonController(createLesson *application.CreateLesson) *CreateLessonController {
	return &CreateLessonController{createLesson: createLesson}
}

func (cl *CreateLessonController) Execute(c *gin.Context) {
	fmt.Println("=== INICIO CREATE LESSON ===")

	moduleID := c.PostForm("moduleId")
	title := c.PostForm("title")
	contentType := c.PostForm("contentType")
	bodyText := c.PostForm("bodyText")
	durationMinutes := c.PostForm("durationMinutes")
	orderIndex := c.PostForm("orderIndex")
	isPreview := c.PostForm("isPreview")

	if moduleID == "" || title == "" || contentType == "" || orderIndex == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obligatorios: moduleId, title, contentType, orderIndex"})
		return
	}

	moduleIDInt, err := parseToInt(moduleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "moduleId debe ser un número"})
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
		ModuleID:        moduleIDInt,
		Title:           title,
		ContentType:     contentType,
		BodyText:        bodyTextPtr,
		DurationMinutes: durationMinutesInt,
		OrderIndex:      orderIndexInt,
		IsPreview:       isPreviewBool,
	}

	// Si es un video, procesar la subida a Bunny.net
	if contentType == "video" {
		fmt.Println("Detectado tipo video, intentando subir a Bunny.net...")
		file, header, err := c.Request.FormFile("video")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere un archivo de video para lecciones tipo 'video'"})
			return
		}
		defer file.Close()

		fmt.Printf("Video recibido: %s, Tamaño: %d bytes\n", header.Filename, header.Size)

		// Leer bytes del video
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("ERROR al leer bytes del video: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar video"})
			return
		}

		// Subir a Bunny.net
		fmt.Println("Subiendo video a Bunny.net...")
		videoID, err := bunny.UploadVideo(fileBytes, header.Filename, title)
		if err != nil {
			fmt.Printf("ERROR al subir a Bunny.net: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al subir video a Bunny.net"})
			return
		}

		fmt.Printf("Video subido exitosamente. Video ID: %s\n", videoID)

		// Guardar IDs de Bunny
		libraryID := os.Getenv("BUNNY_LIBRARY_ID")
		lesson.BunnyLibraryID = &libraryID
		lesson.BunnyVideoID = &videoID
	}

	fmt.Println("Guardando lección en BD...")
	savedLesson, err := cl.createLesson.Execute(lesson)
	if err != nil {
		fmt.Printf("ERROR al guardar lección: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Lección creada exitosamente")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Lección creada exitosamente",
		"lesson": dto.LessonResponse{
			ID:              savedLesson.ID,
			ModuleID:        savedLesson.ModuleID,
			Title:           savedLesson.Title,
			ContentType:     savedLesson.ContentType,
			VideoURL:        savedLesson.GetVideoURL(),
			BodyText:        savedLesson.BodyText,
			DurationMinutes: savedLesson.DurationMinutes,
			OrderIndex:      savedLesson.OrderIndex,
			IsPreview:       savedLesson.IsPreview,
			CreatedAt:       savedLesson.CreatedAt,
		},
	})
}

func parseToInt(s string) (int, error) {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
}
