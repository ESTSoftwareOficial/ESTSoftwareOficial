package adapters

import (
	"database/sql"
	"errors"
	"fmt"

	"estsoftwareoficial/src/resource_types/domain/entities"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(resourceType *entities.ResourceType) (*entities.ResourceType, error) {
	query := `
		INSERT INTO resource_types (name, icon_url) 
		VALUES ($1, $2) 
		RETURNING id
	`

	var id int
	err := ps.conn.QueryRow(query, resourceType.Name, resourceType.IconURL).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("error al guardar tipo de recurso: %v", err)
	}

	resourceType.ID = id
	return resourceType, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.ResourceType, error) {
	query := `
		SELECT id, name, icon_url 
		FROM resource_types 
		WHERE id = $1
	`

	var resourceType entities.ResourceType
	err := ps.conn.QueryRow(query, id).Scan(
		&resourceType.ID,
		&resourceType.Name,
		&resourceType.IconURL,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar tipo de recurso por ID: %v", err)
	}

	return &resourceType, nil
}

func (ps *PostgreSQL) GetByName(name string) (*entities.ResourceType, error) {
	query := `
		SELECT id, name, icon_url 
		FROM resource_types 
		WHERE name = $1
	`

	var resourceType entities.ResourceType
	err := ps.conn.QueryRow(query, name).Scan(
		&resourceType.ID,
		&resourceType.Name,
		&resourceType.IconURL,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar tipo de recurso por nombre: %v", err)
	}

	return &resourceType, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.ResourceType, error) {
	query := `
		SELECT id, name, icon_url 
		FROM resource_types 
		ORDER BY name ASC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todos los tipos de recurso: %v", err)
	}
	defer rows.Close()

	var resourceTypes []*entities.ResourceType
	for rows.Next() {
		var resourceType entities.ResourceType
		err := rows.Scan(
			&resourceType.ID,
			&resourceType.Name,
			&resourceType.IconURL,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear tipo de recurso: %v", err)
		}
		resourceTypes = append(resourceTypes, &resourceType)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar resultados: %v", err)
	}

	return resourceTypes, nil
}

func (ps *PostgreSQL) Update(resourceType *entities.ResourceType) error {
	query := `
		UPDATE resource_types SET 
			name = $2, 
			icon_url = $3
		WHERE id = $1
	`

	result, err := ps.conn.Exec(query, resourceType.ID, resourceType.Name, resourceType.IconURL)
	if err != nil {
		return fmt.Errorf("error al actualizar tipo de recurso: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("tipo de recurso no encontrado")
	}

	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM resource_types WHERE id = $1`

	result, err := ps.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar tipo de recurso: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("tipo de recurso no encontrado")
	}

	return nil
}
