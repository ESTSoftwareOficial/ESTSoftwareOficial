package adapters

import (
	"database/sql"
	"errors"
	"fmt"

	"estsoftwareoficial/src/modules/domain/entities"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(module *entities.Module) (*entities.Module, error) {
	query := `
		INSERT INTO modules (course_id, title, description, order_index) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id
	`

	var id int
	err := ps.conn.QueryRow(
		query,
		module.CourseID,
		module.Title,
		module.Description,
		module.OrderIndex,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("error al guardar módulo: %v", err)
	}

	module.ID = id
	return module, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.Module, error) {
	query := `
		SELECT id, course_id, title, description, order_index 
		FROM modules 
		WHERE id = $1
	`

	var module entities.Module
	err := ps.conn.QueryRow(query, id).Scan(
		&module.ID,
		&module.CourseID,
		&module.Title,
		&module.Description,
		&module.OrderIndex,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar módulo por ID: %v", err)
	}

	return &module, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.Module, error) {
	query := `
		SELECT id, course_id, title, description, order_index 
		FROM modules 
		ORDER BY course_id, order_index ASC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todos los módulos: %v", err)
	}
	defer rows.Close()

	var modules []*entities.Module
	for rows.Next() {
		var module entities.Module
		err := rows.Scan(
			&module.ID,
			&module.CourseID,
			&module.Title,
			&module.Description,
			&module.OrderIndex,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear módulo: %v", err)
		}
		modules = append(modules, &module)
	}

	return modules, nil
}

func (ps *PostgreSQL) GetByCourse(courseID int) ([]*entities.Module, error) {
	query := `
		SELECT id, course_id, title, description, order_index 
		FROM modules 
		WHERE course_id = $1
		ORDER BY order_index ASC
	`

	rows, err := ps.conn.Query(query, courseID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener módulos por curso: %v", err)
	}
	defer rows.Close()

	var modules []*entities.Module
	for rows.Next() {
		var module entities.Module
		err := rows.Scan(
			&module.ID,
			&module.CourseID,
			&module.Title,
			&module.Description,
			&module.OrderIndex,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear módulo: %v", err)
		}
		modules = append(modules, &module)
	}

	return modules, nil
}

func (ps *PostgreSQL) Update(module *entities.Module) error {
	query := `
		UPDATE modules SET 
			title = $2, 
			description = $3, 
			order_index = $4
		WHERE id = $1
	`

	result, err := ps.conn.Exec(
		query,
		module.ID,
		module.Title,
		module.Description,
		module.OrderIndex,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar módulo: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("módulo no encontrado")
	}

	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM modules WHERE id = $1`

	result, err := ps.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar módulo: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("módulo no encontrado")
	}

	return nil
}

func (ps *PostgreSQL) UpdateOrderIndex(id int, orderIndex int) error {
	query := `UPDATE modules SET order_index = $2 WHERE id = $1`

	_, err := ps.conn.Exec(query, id, orderIndex)
	if err != nil {
		return fmt.Errorf("error al actualizar orden del módulo: %v", err)
	}

	return nil
}

func (ps *PostgreSQL) GetTotalModulesByCourse(courseID int) (int, error) {
	query := `SELECT COUNT(*) FROM modules WHERE course_id = $1`

	var total int
	err := ps.conn.QueryRow(query, courseID).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("error al contar módulos: %v", err)
	}

	return total, nil
}
