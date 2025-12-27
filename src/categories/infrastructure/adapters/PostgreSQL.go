package adapters

import (
	"database/sql"
	"errors"
	"fmt"

	"estsoftwareoficial/src/categories/domain/entities"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(category *entities.Category) (*entities.Category, error) {
	query := `
		INSERT INTO categories (name, description, created_at) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`

	var id int
	err := ps.conn.QueryRow(query, category.Name, category.Description, category.CreatedAt).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("error al guardar categoría: %v", err)
	}

	category.ID = id
	return category, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.Category, error) {
	query := `
		SELECT id, name, description, created_at 
		FROM categories 
		WHERE id = $1
	`

	var category entities.Category
	err := ps.conn.QueryRow(query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar categoría por ID: %v", err)
	}

	return &category, nil
}

func (ps *PostgreSQL) GetByName(name string) (*entities.Category, error) {
	query := `
		SELECT id, name, description, created_at 
		FROM categories 
		WHERE name = $1
	`

	var category entities.Category
	err := ps.conn.QueryRow(query, name).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar categoría por nombre: %v", err)
	}

	return &category, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.Category, error) {
	query := `
		SELECT id, name, description, created_at 
		FROM categories 
		ORDER BY name ASC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todas las categorías: %v", err)
	}
	defer rows.Close()

	var categories []*entities.Category
	for rows.Next() {
		var category entities.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear categoría: %v", err)
		}
		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar resultados: %v", err)
	}

	return categories, nil
}

func (ps *PostgreSQL) Update(category *entities.Category) error {
	query := `
		UPDATE categories SET 
			name = $2, 
			description = $3
		WHERE id = $1
	`

	result, err := ps.conn.Exec(query, category.ID, category.Name, category.Description)
	if err != nil {
		return fmt.Errorf("error al actualizar categoría: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("categoría no encontrada")
	}

	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`

	result, err := ps.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar categoría: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("categoría no encontrada")
	}

	return nil
}
