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
			module_id, title, content_type, content_url, body_text, 
			duration_minutes, order_index, is_preview, created_at
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		RETURNING id, created_at
	`

	now := time.Now()
	err := ps.conn.QueryRow(
		query,
		lesson.ModuleID,
		lesson.Title,
		lesson.ContentType,
		lesson.ContentURL,
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
		SELECT id, module_id, title, content_type, content_url, body_text, 
		       duration_minutes, order_index, is_preview, created_at
		FROM lessons 
		WHERE id = $1
	`

	var lesson entities.Lesson
	err := ps.conn.QueryRow(query, id).Scan(
		&lesson.ID,
		&lesson.ModuleID,
		&lesson.Title,
		&lesson.ContentType,
		&lesson.ContentURL,
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
		SELECT id, module_id, title, content_type, content_url, body_text, 
		       duration_minutes, order_index, is_preview, created_at
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
			&lesson.ContentURL,
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
		SELECT id, module_id, title, content_type, content_url, body_text, 
		       duration_minutes, order_index, is_preview, created_at
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
			&lesson.ContentURL,
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
			content_url = $4, 
			body_text = $5, 
			duration_minutes = $6, 
			order_index = $7, 
			is_preview = $8
		WHERE id = $1
	`

	result, err := ps.conn.Exec(
		query,
		lesson.ID,
		lesson.Title,
		lesson.ContentType,
		lesson.ContentURL,
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
