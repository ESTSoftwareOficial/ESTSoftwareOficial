package adapters

import (
	"database/sql"
	"errors"
	"fmt"

	"estsoftwareoficial/src/lesson_resources/domain/entities"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(lessonResource *entities.LessonResource) (*entities.LessonResource, error) {
	query := `
		INSERT INTO lesson_resources (lesson_id, resource_type_id, url, title) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id
	`

	var id int
	err := ps.conn.QueryRow(
		query,
		lessonResource.LessonID,
		lessonResource.ResourceTypeID,
		lessonResource.URL,
		lessonResource.Title,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("error al guardar recurso de lección: %v", err)
	}

	lessonResource.ID = id
	return lessonResource, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.LessonResource, error) {
	query := `
		SELECT id, lesson_id, resource_type_id, url, title 
		FROM lesson_resources 
		WHERE id = $1
	`

	var lessonResource entities.LessonResource
	err := ps.conn.QueryRow(query, id).Scan(
		&lessonResource.ID,
		&lessonResource.LessonID,
		&lessonResource.ResourceTypeID,
		&lessonResource.URL,
		&lessonResource.Title,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar recurso de lección por ID: %v", err)
	}

	return &lessonResource, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.LessonResource, error) {
	query := `
		SELECT id, lesson_id, resource_type_id, url, title 
		FROM lesson_resources 
		ORDER BY lesson_id, id ASC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todos los recursos de lecciones: %v", err)
	}
	defer rows.Close()

	var lessonResources []*entities.LessonResource
	for rows.Next() {
		var lessonResource entities.LessonResource
		err := rows.Scan(
			&lessonResource.ID,
			&lessonResource.LessonID,
			&lessonResource.ResourceTypeID,
			&lessonResource.URL,
			&lessonResource.Title,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear recurso de lección: %v", err)
		}
		lessonResources = append(lessonResources, &lessonResource)
	}

	return lessonResources, nil
}

func (ps *PostgreSQL) GetByLesson(lessonID int) ([]*entities.LessonResource, error) {
	query := `
		SELECT id, lesson_id, resource_type_id, url, title 
		FROM lesson_resources 
		WHERE lesson_id = $1
		ORDER BY id ASC
	`

	rows, err := ps.conn.Query(query, lessonID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener recursos por lección: %v", err)
	}
	defer rows.Close()

	var lessonResources []*entities.LessonResource
	for rows.Next() {
		var lessonResource entities.LessonResource
		err := rows.Scan(
			&lessonResource.ID,
			&lessonResource.LessonID,
			&lessonResource.ResourceTypeID,
			&lessonResource.URL,
			&lessonResource.Title,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear recurso de lección: %v", err)
		}
		lessonResources = append(lessonResources, &lessonResource)
	}

	return lessonResources, nil
}

func (ps *PostgreSQL) Update(lessonResource *entities.LessonResource) error {
	query := `
		UPDATE lesson_resources SET 
			resource_type_id = $2, 
			url = $3, 
			title = $4
		WHERE id = $1
	`

	result, err := ps.conn.Exec(
		query,
		lessonResource.ID,
		lessonResource.ResourceTypeID,
		lessonResource.URL,
		lessonResource.Title,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar recurso de lección: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("recurso de lección no encontrado")
	}

	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM lesson_resources WHERE id = $1`

	result, err := ps.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar recurso de lección: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("recurso de lección no encontrado")
	}

	return nil
}
