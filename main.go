package main

import (
	"database/sql"
	"fmt"
	"log"
	"todo/db_work"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("", "") // your db here
	if err != nil {
		log.Fatal("Ошибка: не удалось подключиться к базе данных")
	}
	defer db.Close()

	// login to user
	user_id, err := db_work.Login(db, "Oleg", "Biba123")
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	// add a note
	note_id, err := db_work.AddNote(db, "написать круд", "эээ ну короче заметки")
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	// connect user and his notes
	err = db_work.Connect(db, user_id, note_id)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	// update note
	err = db_work.UpdateNote(db, "стандоф", "тест на роналда", user_id, 4)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	// show all notes by certain user
	notes, err := db_work.SelectNotesByUser(db, user_id)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	fmt.Println(notes)
}
