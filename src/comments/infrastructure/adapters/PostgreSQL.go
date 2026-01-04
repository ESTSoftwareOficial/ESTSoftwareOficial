package adapters

import (
	"database/sql"
	"errors"
	"estsoftwareoficial/src/comments/domain/entities"
	"fmt"
	"time"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Create(comment *entities.Comment) (*entities.Comment, error) {
	query := `
        INSERT INTO lesson_comments (lesson_id, user_id, comment, created_at, updated_at, is_edited, is_deleted)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `

	now := time.Now()
	comment.CreatedAt = now
	comment.UpdatedAt = now
	comment.IsEdited = false
	comment.IsDeleted = false

	err := ps.conn.QueryRow(
		query,
		comment.LessonID,
		comment.UserID,
		comment.Comment,
		comment.CreatedAt,
		comment.UpdatedAt,
		comment.IsEdited,
		comment.IsDeleted,
	).Scan(&comment.ID)

	if err != nil {
		return nil, fmt.Errorf("error al crear comentario: %v", err)
	}

	return comment, nil
}

func (ps *PostgreSQL) GetByLessonID(lessonID, userID, limit, offset int) ([]*entities.Comment, error) {
	query := `
        SELECT id, lesson_id, user_id, comment, created_at, updated_at, is_edited, is_deleted
        FROM lesson_comments
        WHERE lesson_id = $1 AND is_deleted = false
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `

	rows, err := ps.conn.Query(query, lessonID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error al obtener comentarios: %v", err)
	}
	defer rows.Close()

	var comments []*entities.Comment
	for rows.Next() {
		var comment entities.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.LessonID,
			&comment.UserID,
			&comment.Comment,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.IsEdited,
			&comment.IsDeleted,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear comentario: %v", err)
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (ps *PostgreSQL) GetByID(commentID int) (*entities.Comment, error) {
	query := `
        SELECT id, lesson_id, user_id, comment, created_at, updated_at, is_edited, is_deleted
        FROM lesson_comments
        WHERE id = $1
    `

	var comment entities.Comment
	err := ps.conn.QueryRow(query, commentID).Scan(
		&comment.ID,
		&comment.LessonID,
		&comment.UserID,
		&comment.Comment,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&comment.IsEdited,
		&comment.IsDeleted,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al obtener comentario: %v", err)
	}

	return &comment, nil
}

func (ps *PostgreSQL) Update(comment *entities.Comment) error {
	query := `
        UPDATE lesson_comments
        SET comment = $2, updated_at = $3, is_edited = $4
        WHERE id = $1 AND is_deleted = false
    `

	comment.UpdatedAt = time.Now()
	comment.IsEdited = true

	result, err := ps.conn.Exec(query, comment.ID, comment.Comment, comment.UpdatedAt, comment.IsEdited)
	if err != nil {
		return fmt.Errorf("error al actualizar comentario: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("comentario no encontrado o ya fue eliminado")
	}

	return nil
}

func (ps *PostgreSQL) Delete(commentID int) error {
	query := `
        UPDATE lesson_comments
        SET is_deleted = true
        WHERE id = $1
    `

	result, err := ps.conn.Exec(query, commentID)
	if err != nil {
		return fmt.Errorf("error al eliminar comentario: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("comentario no encontrado")
	}

	return nil
}

func (ps *PostgreSQL) GetTotalByLesson(lessonID int) (int, error) {
	query := `SELECT COUNT(*) FROM lesson_comments WHERE lesson_id = $1 AND is_deleted = false`

	var total int
	err := ps.conn.QueryRow(query, lessonID).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("error al obtener total de comentarios: %v", err)
	}

	return total, nil
}

func (ps *PostgreSQL) AddLike(commentID, userID int) error {
	query := `
        INSERT INTO comment_likes (comment_id, user_id, created_at)
        VALUES ($1, $2, $3)
        ON CONFLICT (comment_id, user_id) DO NOTHING
    `

	_, err := ps.conn.Exec(query, commentID, userID, time.Now())
	if err != nil {
		return fmt.Errorf("error al agregar like: %v", err)
	}

	return nil
}

func (ps *PostgreSQL) RemoveLike(commentID, userID int) error {
	query := `DELETE FROM comment_likes WHERE comment_id = $1 AND user_id = $2`

	_, err := ps.conn.Exec(query, commentID, userID)
	if err != nil {
		return fmt.Errorf("error al quitar like: %v", err)
	}

	return nil
}

func (ps *PostgreSQL) GetLikesCount(commentID int) (int, error) {
	query := `SELECT COUNT(*) FROM comment_likes WHERE comment_id = $1`

	var count int
	err := ps.conn.QueryRow(query, commentID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error al obtener conteo de likes: %v", err)
	}

	return count, nil
}

func (ps *PostgreSQL) UserHasLiked(commentID, userID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM comment_likes WHERE comment_id = $1 AND user_id = $2)`

	var exists bool
	err := ps.conn.QueryRow(query, commentID, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error al verificar like del usuario: %v", err)
	}

	return exists, nil
}

func (ps *PostgreSQL) GetUserInfo(userID int) (string, string, *string, error) {
	query := `SELECT first_name, last_name, profile_photo FROM users WHERE id = $1`

	var firstName, lastName string
	var profilePhoto *string
	err := ps.conn.QueryRow(query, userID).Scan(&firstName, &lastName, &profilePhoto)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", nil, errors.New("usuario no encontrado")
		}
		return "", "", nil, fmt.Errorf("error al obtener información del usuario: %v", err)
	}

	return firstName, lastName, profilePhoto, nil
}
