package db_work

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

func checkForSymbols(symbols string, pass string) bool {
	matched, _ := regexp.MatchString(symbols, pass)
	return matched
}

func checkPassword(password string) error {
	switch {
	case len(password) < 4:
		return fmt.Errorf("пароль должен содержать не менее 4 символов")
	case !checkForSymbols(`[a-z]`, password):
		return fmt.Errorf("пароль должен содержать хотя бы одну строчную букву")
	case !checkForSymbols(`[A-Z]`, password):
		return fmt.Errorf("пароль должен содержать хотя бы одну заглавную букву")
	case !checkForSymbols(`[0-9]`, password):
		return fmt.Errorf("пароль должен содержать хотя бы одну цифру")
	default:
		return nil
	}
}

func CreateNewUser(db *sql.DB, user_name, password string) error {
	createQuery := `CREATE TABLE IF NOT EXISTS users
				(
					user_id serial PRIMARY KEY,
					name varchar(32) NOT NULL UNIQUE,
					password varchar(32) NOT NULL UNIQUE
				);
	`
	_, err := db.Exec(createQuery)
	if err != nil {
		return fmt.Errorf("создание таблицы")
	}

	if len(user_name) == 0 {
		return fmt.Errorf("имя не может быть пустым")
	}

	err = checkPassword(password)
	if err != nil {
		return err
	}
	insertQuery := `
		INSERT INTO users (name, password)
		VALUES ($1, $2);
	`
	if _, err := db.Exec(insertQuery, user_name, password); err != nil {
		return fmt.Errorf("авторизация: %v", err)
	}

	fmt.Println("Успешная авторизация")
	return nil
}
