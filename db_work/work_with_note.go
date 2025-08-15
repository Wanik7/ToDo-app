package db_work

import (
	"database/sql"
	"fmt"
)

type Note struct {
	Title   string
	Content string
}

func AddNote(db *sql.DB, note_title, note_body string) (int, error) {
	createQuery := `CREATE TABLE IF NOT EXISTS notes
				(
					note_id serial PRIMARY KEY,
					title varchar(32) NOT NULL,
					body varchar
				);
	`
	_, err := db.Exec(createQuery)
	if err != nil {
		return 0, fmt.Errorf("создание таблицы")
	}

	if len(note_title) == 0 {
		return 0, fmt.Errorf("название заметки не может быть пустым")
	}

	var note_id int
	insertQuery := `
				INSERT INTO notes (title, body)
				VALUES ($1, $2)
				RETURNING note_id
				`
	if err := db.QueryRow(insertQuery, note_title, note_body).Scan(&note_id); err != nil {
		return 0, fmt.Errorf("авторизация: %v", err)
	}

	fmt.Println("Заметка создана")
	return note_id, nil
}

func SelectNotesByUser(db *sql.DB, user_id int) ([]Note, error) {
	query := `
		SELECT title, body
		FROM operations
		JOIN users USING(user_id)
		JOIN notes USING(note_id)
		WHERE user_id = $1
	`
	rows, err := db.Query(query, user_id)
	if err != nil {
		return []Note{}, fmt.Errorf("не удалось загрузить заметки")
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.Title, &note.Content)
		if err != nil {
			return []Note{}, fmt.Errorf("чтение заметок")
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func DeleteNote(db *sql.DB, user_id, note_id int) error {
	query := `
        DELETE FROM notes 
        USING operations
        WHERE notes.note_id = operations.note_id
        AND operations.user_id = $1
        AND notes.note_id = $2
    `
	_, err := db.Exec(query, user_id, note_id)
	if err != nil {
		return fmt.Errorf("не удалось удалить заметку")
	}

	fmt.Println("Заметка была удалена")
	return nil
}

func UpdateNote(db *sql.DB, newTitle string, newBody string, user_id, note_id int) error {
	query := `
        UPDATE notes
        SET title = $1, body = $2
        WHERE note_id IN (
            SELECT operations.note_id FROM operations
            WHERE operations.user_id = $3 AND operations.note_id = $4
        )
    `

	_, err := db.Exec(query, newTitle, newBody, user_id, note_id)
	if err != nil {
		return fmt.Errorf("не удалось изменить заметку")
	}

	fmt.Println("Заметка была изменена")
	return nil
}
