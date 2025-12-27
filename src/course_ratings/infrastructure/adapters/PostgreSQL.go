package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"estsoftwareoficial/src/course_ratings/domain/entities"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(courseRating *entities.CourseRating) (*entities.CourseRating, error) {
	query := `
		INSERT INTO course_ratings (course_id, user_id, rating, review, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, created_at, updated_at
	`

	now := time.Now()
	err := ps.conn.QueryRow(
		query,
		courseRating.CourseID,
		courseRating.UserID,
		courseRating.Rating,
		courseRating.Review,
		now,
		now,
	).Scan(&courseRating.ID, &courseRating.CreatedAt, &courseRating.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("error al guardar calificación: %v", err)
	}

	return courseRating, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.CourseRating, error) {
	query := `
		SELECT id, course_id, user_id, rating, review, created_at, updated_at 
		FROM course_ratings 
		WHERE id = $1
	`

	var courseRating entities.CourseRating
	err := ps.conn.QueryRow(query, id).Scan(
		&courseRating.ID,
		&courseRating.CourseID,
		&courseRating.UserID,
		&courseRating.Rating,
		&courseRating.Review,
		&courseRating.CreatedAt,
		&courseRating.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar calificación por ID: %v", err)
	}

	return &courseRating, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.CourseRating, error) {
	query := `
		SELECT id, course_id, user_id, rating, review, created_at, updated_at 
		FROM course_ratings 
		ORDER BY created_at DESC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todas las calificaciones: %v", err)
	}
	defer rows.Close()

	var courseRatings []*entities.CourseRating
	for rows.Next() {
		var courseRating entities.CourseRating
		err := rows.Scan(
			&courseRating.ID,
			&courseRating.CourseID,
			&courseRating.UserID,
			&courseRating.Rating,
			&courseRating.Review,
			&courseRating.CreatedAt,
			&courseRating.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear calificación: %v", err)
		}
		courseRatings = append(courseRatings, &courseRating)
	}

	return courseRatings, nil
}

func (ps *PostgreSQL) GetByCourse(courseID int) ([]*entities.CourseRating, error) {
	query := `
		SELECT id, course_id, user_id, rating, review, created_at, updated_at 
		FROM course_ratings 
		WHERE course_id = $1
		ORDER BY created_at DESC
	`

	rows, err := ps.conn.Query(query, courseID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener calificaciones por curso: %v", err)
	}
	defer rows.Close()

	var courseRatings []*entities.CourseRating
	for rows.Next() {
		var courseRating entities.CourseRating
		err := rows.Scan(
			&courseRating.ID,
			&courseRating.CourseID,
			&courseRating.UserID,
			&courseRating.Rating,
			&courseRating.Review,
			&courseRating.CreatedAt,
			&courseRating.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear calificación: %v", err)
		}
		courseRatings = append(courseRatings, &courseRating)
	}

	return courseRatings, nil
}

func (ps *PostgreSQL) GetByUserAndCourse(userID int, courseID int) (*entities.CourseRating, error) {
	query := `
		SELECT id, course_id, user_id, rating, review, created_at, updated_at 
		FROM course_ratings 
		WHERE user_id = $1 AND course_id = $2
	`

	var courseRating entities.CourseRating
	err := ps.conn.QueryRow(query, userID, courseID).Scan(
		&courseRating.ID,
		&courseRating.CourseID,
		&courseRating.UserID,
		&courseRating.Rating,
		&courseRating.Review,
		&courseRating.CreatedAt,
		&courseRating.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar calificación por usuario y curso: %v", err)
	}

	return &courseRating, nil
}

func (ps *PostgreSQL) Update(courseRating *entities.CourseRating) error {
	query := `
		UPDATE course_ratings SET 
			rating = $2, 
			review = $3, 
			updated_at = $4
		WHERE id = $1
	`

	result, err := ps.conn.Exec(
		query,
		courseRating.ID,
		courseRating.Rating,
		courseRating.Review,
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("error al actualizar calificación: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("calificación no encontrada")
	}

	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM course_ratings WHERE id = $1`

	result, err := ps.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar calificación: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("calificación no encontrada")
	}

	return nil
}

func (ps *PostgreSQL) CalculateAverageRating(courseID int) (float64, int, error) {
	query := `
		SELECT COALESCE(AVG(rating), 0) as average, COUNT(*) as total
		FROM course_ratings 
		WHERE course_id = $1
	`

	var average float64
	var total int

	err := ps.conn.QueryRow(query, courseID).Scan(&average, &total)
	if err != nil {
		return 0, 0, fmt.Errorf("error al calcular promedio de calificaciones: %v", err)
	}

	return average, total, nil
}
