package dto

// InstructorDetailDTO representa la información del instructor
type InstructorDetailDTO struct {
	Name   string              `json:"name"`
	Puesto string              `json:"puesto"`
	Redes  []SocialNetworkDTO  `json:"redes"`
}

// SocialNetworkDTO representa una red social
type SocialNetworkDTO struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	URL  string `json:"url"`
}

// ResourceDetailDTO representa un recurso de la lección
type ResourceDetailDTO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	IconURL string `json:"iconUrl"`
}

// LessonDetailResponse es la respuesta del endpoint de detalle
type LessonDetailResponse struct {
	ID               int                   `json:"id"`
	Titulo           string                `json:"titulo"`
	VideoURL         string                `json:"videoUrl"`
	Description      string                `json:"description"`
	Likes            int                   `json:"likes"`
	Instructor       InstructorDetailDTO   `json:"instructor"`
	Recursos         []ResourceDetailDTO   `json:"recursos"`
	ComentariosCount int                   `json:"comentariosCount"`
}