package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"estsoftwareoficial/src/courses/domain/entities"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(course *entities.Course) (*entities.Course, error) {
	query := `
		INSERT INTO courses (
			name_course, description, technology_id, instructor_id, 
			category_id, level, image_url, total_modules, 
			average_rating, total_ratings, duration_hours, 
			created_at, updated_at, is_active
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) 
		RETURNING id, created_at, updated_at
	`

	now := time.Now()
	err := ps.conn.QueryRow(
		query,
		course.NameCourse,
		course.Description,
		course.TechnologyID,
		course.InstructorID,
		course.CategoryID,
		course.Level,
		course.ImageURL,
		0,    // total_modules empieza en 0
		0.00, // average_rating empieza en 0
		0,    // total_ratings empieza en 0
		course.DurationHours,
		now,
		now,
		true, // is_active por defecto true
	).Scan(&course.ID, &course.CreatedAt, &course.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("error al guardar curso: %v", err)
	}

	course.TotalModules = 0
	course.AverageRating = 0.00
	course.TotalRatings = 0
	course.IsActive = true

	return course, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.Course, error) {
	query := `
		SELECT id, name_course, description, technology_id, instructor_id, 
		       category_id, level, image_url, total_modules, average_rating, 
		       total_ratings, duration_hours, created_at, updated_at, is_active
		FROM courses 
		WHERE id = $1
	`

	var course entities.Course
	err := ps.conn.QueryRow(query, id).Scan(
		&course.ID,
		&course.NameCourse,
		&course.Description,
		&course.TechnologyID,
		&course.InstructorID,
		&course.CategoryID,
		&course.Level,
		&course.ImageURL,
		&course.TotalModules,
		&course.AverageRating,
		&course.TotalRatings,
		&course.DurationHours,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar curso por ID: %v", err)
	}

	return &course, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.Course, error) {
	query := `
		SELECT id, name_course, description, technology_id, instructor_id, 
		       category_id, level, image_url, total_modules, average_rating, 
		       total_ratings, duration_hours, created_at, updated_at, is_active
		FROM courses 
		WHERE is_active = true
		ORDER BY created_at DESC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todos los cursos: %v", err)
	}
	defer rows.Close()

	var courses []*entities.Course
	for rows.Next() {
		var course entities.Course
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.TechnologyID,
			&course.InstructorID,
			&course.CategoryID,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

func (ps *PostgreSQL) GetByInstructor(instructorID int) ([]*entities.Course, error) {
	query := `
		SELECT id, name_course, description, technology_id, instructor_id, 
		       category_id, level, image_url, total_modules, average_rating, 
		       total_ratings, duration_hours, created_at, updated_at, is_active
		FROM courses 
		WHERE instructor_id = $1
		ORDER BY created_at DESC
	`

	rows, err := ps.conn.Query(query, instructorID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cursos por instructor: %v", err)
	}
	defer rows.Close()

	var courses []*entities.Course
	for rows.Next() {
		var course entities.Course
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.TechnologyID,
			&course.InstructorID,
			&course.CategoryID,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

func (ps *PostgreSQL) GetByCategory(categoryID int) ([]*entities.Course, error) {
	query := `
		SELECT id, name_course, description, technology_id, instructor_id, 
		       category_id, level, image_url, total_modules, average_rating, 
		       total_ratings, duration_hours, created_at, updated_at, is_active
		FROM courses 
		WHERE category_id = $1 AND is_active = true
		ORDER BY average_rating DESC, created_at DESC
	`

	rows, err := ps.conn.Query(query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cursos por categoría: %v", err)
	}
	defer rows.Close()

	var courses []*entities.Course
	for rows.Next() {
		var course entities.Course
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.TechnologyID,
			&course.InstructorID,
			&course.CategoryID,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

func (ps *PostgreSQL) GetByTechnology(technologyID int) ([]*entities.Course, error) {
	query := `
		SELECT id, name_course, description, technology_id, instructor_id, 
		       category_id, level, image_url, total_modules, average_rating, 
		       total_ratings, duration_hours, created_at, updated_at, is_active
		FROM courses 
		WHERE technology_id = $1 AND is_active = true
		ORDER BY average_rating DESC, created_at DESC
	`

	rows, err := ps.conn.Query(query, technologyID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cursos por tecnología: %v", err)
	}
	defer rows.Close()

	var courses []*entities.Course
	for rows.Next() {
		var course entities.Course
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.TechnologyID,
			&course.InstructorID,
			&course.CategoryID,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

func (ps *PostgreSQL) Update(course *entities.Course) error {
	query := `
		UPDATE courses SET 
			name_course = $2, 
			description = $3, 
			technology_id = $4, 
			category_id = $5, 
			level = $6, 
			image_url = $7, 
			duration_hours = $8, 
			updated_at = $9,
			is_active = $10
		WHERE id = $1
	`

	result, err := ps.conn.Exec(
		query,
		course.ID,
		course.NameCourse,
		course.Description,
		course.TechnologyID,
		course.CategoryID,
		course.Level,
		course.ImageURL,
		course.DurationHours,
		time.Now(),
		course.IsActive,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar curso: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("curso no encontrado")
	}

	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM courses WHERE id = $1`

	result, err := ps.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar curso: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("curso no encontrado")
	}

	return nil
}

func (ps *PostgreSQL) UpdateTotalModules(courseID int, totalModules int) error {
	query := `
		UPDATE courses 
		SET total_modules = $2, updated_at = $3
		WHERE id = $1
	`

	_, err := ps.conn.Exec(query, courseID, totalModules, time.Now())
	if err != nil {
		return fmt.Errorf("error al actualizar total de módulos: %v", err)
	}

	return nil
}

func (ps *PostgreSQL) UpdateRating(courseID int, averageRating float64, totalRatings int) error {
	query := `
		UPDATE courses 
		SET average_rating = $2, total_ratings = $3, updated_at = $4
		WHERE id = $1
	`

	_, err := ps.conn.Exec(query, courseID, averageRating, totalRatings, time.Now())
	if err != nil {
		return fmt.Errorf("error al actualizar rating del curso: %v", err)
	}

	return nil
}

func (ps *PostgreSQL) Search(keyword string, categoryID *int, technologyID *int, level *string, minRating *float64) ([]*entities.Course, error) {
	query := `
		SELECT id, name_course, description, technology_id, instructor_id, 
		       category_id, level, image_url, total_modules, average_rating, 
		       total_ratings, duration_hours, created_at, updated_at, is_active
		FROM courses 
		WHERE is_active = true
	`

	args := []interface{}{}
	argPosition := 1

	// Búsqueda por palabra clave (en nombre o descripción)
	if keyword != "" {
		query += fmt.Sprintf(" AND (LOWER(name_course) LIKE $%d OR LOWER(description) LIKE $%d)", argPosition, argPosition)
		args = append(args, "%"+strings.ToLower(keyword)+"%")
		argPosition++
	}

	// Filtro por categoría
	if categoryID != nil {
		query += fmt.Sprintf(" AND category_id = $%d", argPosition)
		args = append(args, *categoryID)
		argPosition++
	}

	// Filtro por tecnología
	if technologyID != nil {
		query += fmt.Sprintf(" AND technology_id = $%d", argPosition)
		args = append(args, *technologyID)
		argPosition++
	}

	// Filtro por nivel
	if level != nil && *level != "" {
		query += fmt.Sprintf(" AND level = $%d", argPosition)
		args = append(args, *level)
		argPosition++
	}

	// Filtro por calificación mínima
	if minRating != nil {
		query += fmt.Sprintf(" AND average_rating >= $%d", argPosition)
		args = append(args, *minRating)
	}

	query += " ORDER BY average_rating DESC, created_at DESC"

	rows, err := ps.conn.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error al buscar cursos: %v", err)
	}
	defer rows.Close()

	var courses []*entities.Course
	for rows.Next() {
		var course entities.Course
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.TechnologyID,
			&course.InstructorID,
			&course.CategoryID,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

// GetByIDWithRelations obtiene un curso por ID con datos relacionados
func (ps *PostgreSQL) GetByIDWithRelations(id int) (*entities.CourseWithRelations, error) {
	query := `
		SELECT 
			c.id, 
			c.name_course, 
			c.description, 
			c.level, 
			c.image_url, 
			c.total_modules, 
			c.average_rating, 
			c.total_ratings, 
			c.duration_hours, 
			c.created_at, 
			c.updated_at, 
			c.is_active,
			t.name AS technology_name,
			t.icon AS technology_image,
			CONCAT(u.first_name, ' ', u.last_name) AS instructor_name,
			u.profile_photo AS instructor_image,
			cat.name AS category_name
		FROM courses c
		INNER JOIN technologies t ON c.technology_id = t.id
		INNER JOIN users u ON c.instructor_id = u.id
		INNER JOIN categories cat ON c.category_id = cat.id
		WHERE c.id = $1
	`

	var course entities.CourseWithRelations
	err := ps.conn.QueryRow(query, id).Scan(
		&course.ID,
		&course.NameCourse,
		&course.Description,
		&course.Level,
		&course.ImageURL,
		&course.TotalModules,
		&course.AverageRating,
		&course.TotalRatings,
		&course.DurationHours,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.IsActive,
		&course.TechnologyName,
		&course.TechnologyImage,
		&course.InstructorName,
		&course.InstructorImage,
		&course.CategoryName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar curso por ID: %v", err)
	}

	return &course, nil
}

// GetAllWithRelations obtiene todos los cursos con datos relacionados
func (ps *PostgreSQL) GetAllWithRelations() ([]*entities.CourseWithRelations, error) {
	query := `
		SELECT 
			c.id, 
			c.name_course, 
			c.description, 
			c.level, 
			c.image_url, 
			c.total_modules, 
			c.average_rating, 
			c.total_ratings, 
			c.duration_hours, 
			c.created_at, 
			c.updated_at, 
			c.is_active,
			t.name AS technology_name,
			t.icon AS technology_image,
			CONCAT(u.first_name, ' ', u.last_name) AS instructor_name,
			u.profile_photo AS instructor_image,
			cat.name AS category_name
		FROM courses c
		INNER JOIN technologies t ON c.technology_id = t.id
		INNER JOIN users u ON c.instructor_id = u.id
		INNER JOIN categories cat ON c.category_id = cat.id
		WHERE c.is_active = true
		ORDER BY c.created_at DESC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todos los cursos: %v", err)
	}
	defer rows.Close()

	var courses []*entities.CourseWithRelations
	for rows.Next() {
		var course entities.CourseWithRelations
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
			&course.TechnologyName,
			&course.TechnologyImage,
			&course.InstructorName,
			&course.InstructorImage,
			&course.CategoryName,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

// GetByInstructorWithRelations obtiene cursos por instructor con datos relacionados
func (ps *PostgreSQL) GetByInstructorWithRelations(instructorID int) ([]*entities.CourseWithRelations, error) {
	query := `
		SELECT 
			c.id, 
			c.name_course, 
			c.description, 
			c.level, 
			c.image_url, 
			c.total_modules, 
			c.average_rating, 
			c.total_ratings, 
			c.duration_hours, 
			c.created_at, 
			c.updated_at, 
			c.is_active,
			t.name AS technology_name,
			t.icon AS technology_image,
			CONCAT(u.first_name, ' ', u.last_name) AS instructor_name,
			u.profile_photo AS instructor_image,
			cat.name AS category_name
		FROM courses c
		INNER JOIN technologies t ON c.technology_id = t.id
		INNER JOIN users u ON c.instructor_id = u.id
		INNER JOIN categories cat ON c.category_id = cat.id
		WHERE c.instructor_id = $1
		ORDER BY c.created_at DESC
	`

	rows, err := ps.conn.Query(query, instructorID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cursos por instructor: %v", err)
	}
	defer rows.Close()

	var courses []*entities.CourseWithRelations
	for rows.Next() {
		var course entities.CourseWithRelations
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
			&course.TechnologyName,
			&course.TechnologyImage,
			&course.InstructorName,
			&course.InstructorImage,
			&course.CategoryName,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

// GetByCategoryWithRelations obtiene cursos por categoría con datos relacionados
func (ps *PostgreSQL) GetByCategoryWithRelations(categoryID int) ([]*entities.CourseWithRelations, error) {
	query := `
		SELECT 
			c.id, 
			c.name_course, 
			c.description, 
			c.level, 
			c.image_url, 
			c.total_modules, 
			c.average_rating, 
			c.total_ratings, 
			c.duration_hours, 
			c.created_at, 
			c.updated_at, 
			c.is_active,
			t.name AS technology_name,
			t.icon AS technology_image,
			CONCAT(u.first_name, ' ', u.last_name) AS instructor_name,
			u.profile_photo AS instructor_image,
			cat.name AS category_name
		FROM courses c
		INNER JOIN technologies t ON c.technology_id = t.id
		INNER JOIN users u ON c.instructor_id = u.id
		INNER JOIN categories cat ON c.category_id = cat.id
		WHERE c.category_id = $1 AND c.is_active = true
		ORDER BY c.average_rating DESC, c.created_at DESC
	`

	rows, err := ps.conn.Query(query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cursos por categoría: %v", err)
	}
	defer rows.Close()

	var courses []*entities.CourseWithRelations
	for rows.Next() {
		var course entities.CourseWithRelations
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
			&course.TechnologyName,
			&course.TechnologyImage,
			&course.InstructorName,
			&course.InstructorImage,
			&course.CategoryName,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

// GetByTechnologyWithRelations obtiene cursos por tecnología con datos relacionados
func (ps *PostgreSQL) GetByTechnologyWithRelations(technologyID int) ([]*entities.CourseWithRelations, error) {
	query := `
		SELECT 
			c.id, 
			c.name_course, 
			c.description, 
			c.level, 
			c.image_url, 
			c.total_modules, 
			c.average_rating, 
			c.total_ratings, 
			c.duration_hours, 
			c.created_at, 
			c.updated_at, 
			c.is_active,
			t.name AS technology_name,
			t.icon AS technology_image,
			CONCAT(u.first_name, ' ', u.last_name) AS instructor_name,
			u.profile_photo AS instructor_image,
			cat.name AS category_name
		FROM courses c
		INNER JOIN technologies t ON c.technology_id = t.id
		INNER JOIN users u ON c.instructor_id = u.id
		INNER JOIN categories cat ON c.category_id = cat.id
		WHERE c.technology_id = $1 AND c.is_active = true
		ORDER BY c.average_rating DESC, c.created_at DESC
	`

	rows, err := ps.conn.Query(query, technologyID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cursos por tecnología: %v", err)
	}
	defer rows.Close()

	var courses []*entities.CourseWithRelations
	for rows.Next() {
		var course entities.CourseWithRelations
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
			&course.TechnologyName,
			&course.TechnologyImage,
			&course.InstructorName,
			&course.InstructorImage,
			&course.CategoryName,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

// SearchWithRelations busca cursos con filtros y datos relacionados
func (ps *PostgreSQL) SearchWithRelations(keyword string, categoryID *int, technologyID *int, level *string, minRating *float64) ([]*entities.CourseWithRelations, error) {
	query := `
		SELECT 
			c.id, 
			c.name_course, 
			c.description, 
			c.level, 
			c.image_url, 
			c.total_modules, 
			c.average_rating, 
			c.total_ratings, 
			c.duration_hours, 
			c.created_at, 
			c.updated_at, 
			c.is_active,
			t.name AS technology_name,
			t.icon AS technology_image,
			CONCAT(u.first_name, ' ', u.last_name) AS instructor_name,
			u.profile_photo AS instructor_image,
			cat.name AS category_name
		FROM courses c
		INNER JOIN technologies t ON c.technology_id = t.id
		INNER JOIN users u ON c.instructor_id = u.id
		INNER JOIN categories cat ON c.category_id = cat.id
		WHERE c.is_active = true
	`

	args := []interface{}{}
	argPosition := 1

	if keyword != "" {
		query += fmt.Sprintf(" AND (LOWER(c.name_course) LIKE $%d OR LOWER(c.description) LIKE $%d)", argPosition, argPosition)
		args = append(args, "%"+strings.ToLower(keyword)+"%")
		argPosition++
	}

	if categoryID != nil {
		query += fmt.Sprintf(" AND c.category_id = $%d", argPosition)
		args = append(args, *categoryID)
		argPosition++
	}

	if technologyID != nil {
		query += fmt.Sprintf(" AND c.technology_id = $%d", argPosition)
		args = append(args, *technologyID)
		argPosition++
	}

	if level != nil && *level != "" {
		query += fmt.Sprintf(" AND c.level = $%d", argPosition)
		args = append(args, *level)
		argPosition++
	}

	if minRating != nil {
		query += fmt.Sprintf(" AND c.average_rating >= $%d", argPosition)
		args = append(args, *minRating)
	}

	query += " ORDER BY c.average_rating DESC, c.created_at DESC"

	rows, err := ps.conn.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error al buscar cursos: %v", err)
	}
	defer rows.Close()

	var courses []*entities.CourseWithRelations
	for rows.Next() {
		var course entities.CourseWithRelations
		err := rows.Scan(
			&course.ID,
			&course.NameCourse,
			&course.Description,
			&course.Level,
			&course.ImageURL,
			&course.TotalModules,
			&course.AverageRating,
			&course.TotalRatings,
			&course.DurationHours,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
			&course.TechnologyName,
			&course.TechnologyImage,
			&course.InstructorName,
			&course.InstructorImage,
			&course.CategoryName,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear curso: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}