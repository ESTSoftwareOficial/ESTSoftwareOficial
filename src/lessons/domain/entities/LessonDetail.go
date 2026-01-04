package entities

import "time"

// LessonDetail representa el detalle completo de una lección con joins
type LessonDetail struct {
	// Datos de la lección
	ID          int
	Title       string
	VideoURL    string
	Description string
	CreatedAt   time.Time

	// Datos del instructor
	InstructorName     string
	InstructorJobTitle *string
	InstructorPhoto    *string

	// Redes sociales del instructor
	GithubURL    *string
	LinkedinURL  *string
	FacebookURL  *string
	XURL         *string
	YoutubeURL   *string
	InstagramURL *string
	PortfolioURL *string

	// Recursos de la lección
	Resources []LessonResource

	// Contadores
	LikesCount    int
	CommentsCount int
}

// LessonResource representa un recurso de la lección
type LessonResource struct {
	ID      int
	Title   string
	URL     string
	IconURL string
}