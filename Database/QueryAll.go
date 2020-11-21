package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type todo struct {
	ID     int
	Title  string
	Status string
}

func main() {

	db, err := sql.Open("postgres", "postgres://oeenwpjw:zQqtsCaL5VoY3x-NX8dbzfH7unkJ9Lb0@hansken.db.elephantsql.com:5432/oeenwpjw")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("Cannot prepare query all row statement ", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Cannot query all todos ", err)
	}

	var todos = []todo{}
	for rows.Next() {

		var t todo
		var id int
		var title string
		var status string

		err = rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("Cannot scan row into variables ", err)
		}

		t = todo{id, title, status}
		todos = append(todos, t)
		//fmt.Println(id, title, status)
	}

	fmt.Printf("query all todos success = %#v", todos)
}
