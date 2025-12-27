package adapters

import (
	"database/sql"
	"errors"
	"estsoftwareoficial/src/users/domain/entities"
	"fmt"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(user *entities.User) (*entities.User, error) {
	query := `
		INSERT INTO users (
			first_name, second_name, last_name, second_last_name, 
			email, secondary_email, password, profile_photo, registration_date, 
			role_id, oauth_provider, oauth_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
		RETURNING id
	`

	var id int
	err := ps.conn.QueryRow(
		query,
		user.FirstName,
		user.SecondName,
		user.LastName,
		user.SecondLastName,
		user.Email,
		user.SecondaryEmail,
		user.Password,
		user.ProfilePhoto,
		user.RegistrationDate,
		user.RoleID,
		user.OAuthProvider,
		user.OAuthID,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("error al guardar usuario: %v", err)
	}

	user.ID = id
	return user, nil
}

func (ps *PostgreSQL) GetByEmail(email string) (*entities.User, error) {
	query := `
		SELECT id, first_name, second_name, last_name, second_last_name, 
		       email, secondary_email, password, profile_photo, registration_date, 
		       role_id, oauth_provider, oauth_id
		FROM users 
		WHERE email = $1
	`

	var user entities.User
	err := ps.conn.QueryRow(query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.SecondName,
		&user.LastName,
		&user.SecondLastName,
		&user.Email,
		&user.SecondaryEmail,
		&user.Password,
		&user.ProfilePhoto,
		&user.RegistrationDate,
		&user.RoleID,
		&user.OAuthProvider,
		&user.OAuthID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar usuario por email: %v", err)
	}

	return &user, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.User, error) {
	query := `
		SELECT id, first_name, second_name, last_name, second_last_name, 
		       email, secondary_email, password, profile_photo, registration_date, 
		       role_id, oauth_provider, oauth_id
		FROM users 
		WHERE id = $1
	`

	var user entities.User
	err := ps.conn.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.SecondName,
		&user.LastName,
		&user.SecondLastName,
		&user.Email,
		&user.SecondaryEmail,
		&user.Password,
		&user.ProfilePhoto,
		&user.RegistrationDate,
		&user.RoleID,
		&user.OAuthProvider,
		&user.OAuthID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar usuario por ID: %v", err)
	}

	return &user, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.User, error) {
	query := `
		SELECT id, first_name, second_name, last_name, second_last_name, 
		       email, secondary_email, password, profile_photo, registration_date, 
		       role_id, oauth_provider, oauth_id
		FROM users 
		ORDER BY registration_date DESC
	`

	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todos los usuarios: %v", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.SecondName,
			&user.LastName,
			&user.SecondLastName,
			&user.Email,
			&user.SecondaryEmail,
			&user.Password,
			&user.ProfilePhoto,
			&user.RegistrationDate,
			&user.RoleID,
			&user.OAuthProvider,
			&user.OAuthID,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear usuario: %v", err)
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar resultados: %v", err)
	}

	return users, nil
}

func (ps *PostgreSQL) Update(user *entities.User) error {
	query := `
		UPDATE users SET 
			first_name = $2, 
			second_name = $3, 
			last_name = $4, 
			second_last_name = $5, 
			email = $6, 
			secondary_email = $7, 
			profile_photo = $8,
			role_id = $9
		WHERE id = $1
	`

	result, err := ps.conn.Exec(
		query,
		user.ID,
		user.FirstName,
		user.SecondName,
		user.LastName,
		user.SecondLastName,
		user.Email,
		user.SecondaryEmail,
		user.ProfilePhoto,
		user.RoleID,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("usuario no encontrado")
	}

	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := ps.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("usuario no encontrado")
	}

	return nil
}

func (ps *PostgreSQL) GetTotal() (int, error) {
	query := `SELECT COUNT(*) FROM users`

	var total int
	err := ps.conn.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("error al obtener total de usuarios: %v", err)
	}

	return total, nil
}
