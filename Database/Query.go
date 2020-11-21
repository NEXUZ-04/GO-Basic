package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://oeenwpjw:zQqtsCaL5VoY3x-NX8dbzfH7unkJ9Lb0@hansken.db.elephantsql.com:5432/oeenwpjw")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("Cannot prepare query one row statement ", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var id int
	var title string
	var status string

	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("Cannot scan row into variables ", err)
	}

	fmt.Println("One row : ", id, title, status)
}
