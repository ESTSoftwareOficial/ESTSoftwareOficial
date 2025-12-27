package adapters

import (
	"database/sql"
	"errors"
	"fmt"

	"estsoftwareoficial/src/technologies/domain/entities"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(technology *entities.Technology) (*entities.Technology, error) {
	query := `
		INSERT INTO technologies (name, icon) 
		VALUES ($1, $2) 
		RETURNING id
	`

	var id int
	err := ps.conn.QueryRow(query, technology.Name, technology.Icon).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("error al guardar tecnología: %v", err)
	}

	technology.ID = id
	return technology, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.Technology, error) {
	query := `
		SELECT id, name, icon 
		FROM technologies 
		WHERE id = $1
	`

	var technology entities.Technology
	err := ps.conn.QueryRow(query, id).Scan(
		&technology.ID,
		&technology.Name,
		&technology.Icon,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar tecnología por ID: %v", err)
	}

	return &technology, nil
}

func (ps *PostgreSQL) GetByName(name string) (*entities.Technology, error) {
	query := `
		SELECT id, name, icon 
		FROM technologies 
		WHERE name = $1
	`

	var technology entities.Technology
	err := ps.conn.QueryRow(query, name).Scan(
		&technology.ID,
		&technology.Name,
		&technology.Icon,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar tecnología por nombre: %v", err)
	}

	return &technology, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.Technology, error) {
	query := `
		SELECT id, name, icon 
		FROM technologies 
		ORDER BY name ASC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todas las tecnologías: %v", err)
	}
	defer rows.Close()

	var technologies []*entities.Technology
	for rows.Next() {
		var technology entities.Technology
		err := rows.Scan(
			&technology.ID,
			&technology.Name,
			&technology.Icon,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear tecnología: %v", err)
		}
		technologies = append(technologies, &technology)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar resultados: %v", err)
	}

	return technologies, nil
}

func (ps *PostgreSQL) Update(technology *entities.Technology) error {
	query := `
		UPDATE technologies SET 
			name = $2, 
			icon = $3
		WHERE id = $1
	`

	result, err := ps.conn.Exec(query, technology.ID, technology.Name, technology.Icon)
	if err != nil {
		return fmt.Errorf("error al actualizar tecnología: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("tecnología no encontrada")
	}

	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM technologies WHERE id = $1`

	result, err := ps.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar tecnología: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("tecnología no encontrada")
	}

	return nil
}
