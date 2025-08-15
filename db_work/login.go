package db_work

import (
	"database/sql"
	"fmt"
)

func Login(db *sql.DB, user_name, user_password string) (int, error) {
	var user_id int
	query := `
			SELECT user_id
			FROM users
			WHERE name = $1 AND password = $2
			`
	err := db.QueryRow(query, user_name, user_password).Scan(&user_id)
	if err != nil {
		return 0, fmt.Errorf("неверный логин или пароль")
	}

	fmt.Println("Удачная авторизация")
	return user_id, nil
}
