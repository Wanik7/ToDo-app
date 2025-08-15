package db_work

import (
	"database/sql"
	"fmt"
)

func Connect(db *sql.DB, user_id, note_id int) error {
	createQuery := `
			CREATE TABLE IF NOT EXISTS operations
			(
				user_id integer REFERENCES users(user_id) ON DELETE CASCADE,
				note_id integer REFERENCES notes(note_id) ON DELETE CASCADE
			)
	`
	_, err := db.Exec(createQuery)
	if err != nil {
		return fmt.Errorf("создание таблицы")
	}

	insertQuery := `
			INSERT INTO operations
			VALUES ($1, $2)
	`
	_, err = db.Exec(insertQuery, user_id, note_id)
	if err != nil {
		return fmt.Errorf("не удалось создать заметку")
	}

	return nil
}
