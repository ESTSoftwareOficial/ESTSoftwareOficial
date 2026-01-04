package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetLessonDetailController struct {
	getLessonDetail *application.GetLessonDetail
}

func NewGetLessonDetailController(getLessonDetail *application.GetLessonDetail) *GetLessonDetailController {
	return &GetLessonDetailController{getLessonDetail: getLessonDetail}
}

func (gl *GetLessonDetailController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	lessonDetail, err := gl.getLessonDetail.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Construir redes sociales del instructor
	redes := []dto.SocialNetworkDTO{}
	
	if lessonDetail.GithubURL != nil && *lessonDetail.GithubURL != "" {
		redes = append(redes, dto.SocialNetworkDTO{
			Name: "github",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/github/github-original.svg",
			URL:  *lessonDetail.GithubURL,
		})
	}
	
	if lessonDetail.LinkedinURL != nil && *lessonDetail.LinkedinURL != "" {
		redes = append(redes, dto.SocialNetworkDTO{
			Name: "linkedin",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/linkedin/linkedin-original.svg",
			URL:  *lessonDetail.LinkedinURL,
		})
	}

	if lessonDetail.FacebookURL != nil && *lessonDetail.FacebookURL != "" {
		redes = append(redes, dto.SocialNetworkDTO{
			Name: "facebook",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/facebook/facebook-original.svg",
			URL:  *lessonDetail.FacebookURL,
		})
	}

	if lessonDetail.XURL != nil && *lessonDetail.XURL != "" {
		redes = append(redes, dto.SocialNetworkDTO{
			Name: "x",
			Icon: "https://cdn.simpleicons.org/x",
			URL:  *lessonDetail.XURL,
		})
	}

	if lessonDetail.YoutubeURL != nil && *lessonDetail.YoutubeURL != "" {
		redes = append(redes, dto.SocialNetworkDTO{
			Name: "youtube",
			Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/youtube/youtube-original.svg",
			URL:  *lessonDetail.YoutubeURL,
		})
	}

	if lessonDetail.InstagramURL != nil && *lessonDetail.InstagramURL != "" {
		redes = append(redes, dto.SocialNetworkDTO{
			Name: "instagram",
			Icon: "https://cdn.simpleicons.org/instagram",
			URL:  *lessonDetail.InstagramURL,
		})
	}

	if lessonDetail.PortfolioURL != nil && *lessonDetail.PortfolioURL != "" {
		redes = append(redes, dto.SocialNetworkDTO{
			Name: "portfolio",
			Icon: "https://cdn.simpleicons.org/internetarchive",
			URL:  *lessonDetail.PortfolioURL,
		})
	}

	// Construir recursos
	recursos := []dto.ResourceDetailDTO{}
	for _, res := range lessonDetail.Resources {
		recursos = append(recursos, dto.ResourceDetailDTO{
			ID:      res.ID,
			Title:   res.Title,
			URL:     res.URL,
			IconURL: res.IconURL,
		})
	}

	// Construir respuesta
	response := dto.LessonDetailResponse{
		ID:          lessonDetail.ID,
		Titulo:      lessonDetail.Title,
		VideoURL:    lessonDetail.VideoURL,
		Description: lessonDetail.Description,
		Likes:       lessonDetail.LikesCount,
		Instructor: dto.InstructorDetailDTO{
			Name:   lessonDetail.InstructorName,
			Puesto: getJobTitle(lessonDetail.InstructorJobTitle),
			Redes:  redes,
		},
		Recursos:         recursos,
		ComentariosCount: lessonDetail.CommentsCount,
	}

	c.JSON(http.StatusOK, response)
}

func getJobTitle(jobTitle *string) string {
	if jobTitle != nil {
		return *jobTitle
	}
	return "Instructor"
}