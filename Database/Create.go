package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //_ : Unused dependency
)

//Init function will be called even if no one called
func init() {
	fmt.Println("Hello")
}

func main() {
	db, err := sql.Open("postgres", "postgres://oeenwpjw:zQqtsCaL5VoY3x-NX8dbzfH7unkJ9Lb0@hansken.db.elephantsql.com:5432/oeenwpjw")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	defer db.Close()
	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")
}
