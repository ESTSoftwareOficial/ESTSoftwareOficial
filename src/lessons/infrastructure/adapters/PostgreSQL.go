package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"estsoftwareoficial/src/lessons/domain/entities"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(lesson *entities.Lesson) (*entities.Lesson, error) {
	query := `
		INSERT INTO lessons (
			module_id, title, content_type, bunny_library_id, bunny_video_id,
			body_text, duration_minutes, order_index, is_preview, created_at
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		RETURNING id, created_at
	`

	now := time.Now()
	err := ps.conn.QueryRow(
		query,
		lesson.ModuleID,
		lesson.Title,
		lesson.ContentType,
		lesson.BunnyLibraryID,
		lesson.BunnyVideoID,
		lesson.BodyText,
		lesson.DurationMinutes,
		lesson.OrderIndex,
		lesson.IsPreview,
		now,
	).Scan(&lesson.ID, &lesson.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("error al guardar lección: %v", err)
	}

	return lesson, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.Lesson, error) {
	query := `
		SELECT id, module_id, title, content_type, bunny_library_id, bunny_video_id,
		       body_text, duration_minutes, order_index, is_preview, created_at
		FROM lessons 
		WHERE id = $1
	`

	var lesson entities.Lesson
	err := ps.conn.QueryRow(query, id).Scan(
		&lesson.ID,
		&lesson.ModuleID,
		&lesson.Title,
		&lesson.ContentType,
		&lesson.BunnyLibraryID,
		&lesson.BunnyVideoID,
		&lesson.BodyText,
		&lesson.DurationMinutes,
		&lesson.OrderIndex,
		&lesson.IsPreview,
		&lesson.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar lección por ID: %v", err)
	}

	return &lesson, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.Lesson, error) {
	query := `
		SELECT id, module_id, title, content_type, bunny_library_id, bunny_video_id,
		       body_text, duration_minutes, order_index, is_preview, created_at
		FROM lessons 
		ORDER BY module_id, order_index ASC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todas las lecciones: %v", err)
	}
	defer rows.Close()

	var lessons []*entities.Lesson
	for rows.Next() {
		var lesson entities.Lesson
		err := rows.Scan(
			&lesson.ID,
			&lesson.ModuleID,
			&lesson.Title,
			&lesson.ContentType,
			&lesson.BunnyLibraryID,
			&lesson.BunnyVideoID,
			&lesson.BodyText,
			&lesson.DurationMinutes,
			&lesson.OrderIndex,
			&lesson.IsPreview,
			&lesson.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear lección: %v", err)
		}
		lessons = append(lessons, &lesson)
	}

	return lessons, nil
}

func (ps *PostgreSQL) GetByModule(moduleID int) ([]*entities.Lesson, error) {
	query := `
		SELECT id, module_id, title, content_type, bunny_library_id, bunny_video_id,
		       body_text, duration_minutes, order_index, is_preview, created_at
		FROM lessons 
		WHERE module_id = $1
		ORDER BY order_index ASC
	`

	rows, err := ps.conn.Query(query, moduleID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener lecciones por módulo: %v", err)
	}
	defer rows.Close()

	var lessons []*entities.Lesson
	for rows.Next() {
		var lesson entities.Lesson
		err := rows.Scan(
			&lesson.ID,
			&lesson.ModuleID,
			&lesson.Title,
			&lesson.ContentType,
			&lesson.BunnyLibraryID,
			&lesson.BunnyVideoID,
			&lesson.BodyText,
			&lesson.DurationMinutes,
			&lesson.OrderIndex,
			&lesson.IsPreview,
			&lesson.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear lección: %v", err)
		}
		lessons = append(lessons, &lesson)
	}

	return lessons, nil
}

func (ps *PostgreSQL) Update(lesson *entities.Lesson) error {
	query := `
		UPDATE lessons SET 
			title = $2, 
			content_type = $3, 
			bunny_library_id = $4,
			bunny_video_id = $5,
			body_text = $6, 
			duration_minutes = $7, 
			order_index = $8, 
			is_preview = $9
		WHERE id = $1
	`

	result, err := ps.conn.Exec(
		query,
		lesson.ID,
		lesson.Title,
		lesson.ContentType,
		lesson.BunnyLibraryID,
		lesson.BunnyVideoID,
		lesson.BodyText,
		lesson.DurationMinutes,
		lesson.OrderIndex,
		lesson.IsPreview,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar lección: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("lección no encontrada")
	}

	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM lessons WHERE id = $1`

	result, err := ps.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar lección: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("lección no encontrada")
	}

	return nil
}

func (ps *PostgreSQL) UpdateOrderIndex(id int, orderIndex int) error {
	query := `UPDATE lessons SET order_index = $2 WHERE id = $1`

	_, err := ps.conn.Exec(query, id, orderIndex)
	if err != nil {
		return fmt.Errorf("error al actualizar orden de la lección: %v", err)
	}

	return nil
}

// GetDetailByID obtiene el detalle completo de una lección con JOINs
func (ps *PostgreSQL) GetDetailByID(id int) (*entities.LessonDetail, error) {
	// Obtiene datos básicos de la lección + instructor + contadores
	query := `
		SELECT 
			l.id,
			l.title,
			COALESCE(l.bunny_video_id, '') AS video_url,
			COALESCE(l.description, '') AS description,
			l.created_at,
			CONCAT(u.first_name, ' ', u.last_name) AS instructor_name,
			u.job_title,
			u.profile_photo,
			u.github_url,
			u.linkedin_url,
			u.facebook_url,
			u.x_url,
			u.youtube_url,
			u.instagram_url,
			u.portfolio_url,
			COALESCE((SELECT COUNT(*) FROM module_likes ml WHERE ml.module_id = l.module_id), 0) AS likes_count,
			COALESCE((SELECT COUNT(*) FROM lesson_comments lc WHERE lc.lesson_id = l.id), 0) AS comments_count
		FROM lessons l
		INNER JOIN modules m ON l.module_id = m.id
		INNER JOIN courses c ON m.course_id = c.id
		INNER JOIN users u ON c.instructor_id = u.id
		WHERE l.id = $1
	`

	var lesson entities.LessonDetail
	err := ps.conn.QueryRow(query, id).Scan(
		&lesson.ID,
		&lesson.Title,
		&lesson.VideoURL,
		&lesson.Description,
		&lesson.CreatedAt,
		&lesson.InstructorName,
		&lesson.InstructorJobTitle,
		&lesson.InstructorPhoto,
		&lesson.GithubURL,
		&lesson.LinkedinURL,
		&lesson.FacebookURL,
		&lesson.XURL,
		&lesson.YoutubeURL,
		&lesson.InstagramURL,
		&lesson.PortfolioURL,
		&lesson.LikesCount,
		&lesson.CommentsCount,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("lección no encontrada")
		}
		return nil, fmt.Errorf("error al obtener lección: %v", err)
	}

	// Obtener recursos de la lección
	resourcesQuery := `
		SELECT 
			lr.id,
			COALESCE(lr.title, rt.name) AS title,
			lr.url,
			COALESCE(rt.icon_url, '') AS icon_url
		FROM lesson_resources lr
		INNER JOIN resource_types rt ON lr.resource_type_id = rt.id
		WHERE lr.lesson_id = $1
		ORDER BY lr.id
	`

	rows, err := ps.conn.Query(resourcesQuery, id)
	if err != nil {
		return nil, fmt.Errorf("error al obtener recursos: %v", err)
	}
	defer rows.Close()

	lesson.Resources = []entities.LessonResource{}
	for rows.Next() {
		var resource entities.LessonResource
		err := rows.Scan(
			&resource.ID,
			&resource.Title,
			&resource.URL,
			&resource.IconURL,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear recurso: %v", err)
		}
		lesson.Resources = append(lesson.Resources, resource)
	}

	return &lesson, nil
}