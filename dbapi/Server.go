package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type DB struct {
	db *sql.DB
}

var Database DB

func (d *DB) Connect() {
	var err error
	d.db, err = sql.Open("postgres", "postgres://oeenwpjw:zQqtsCaL5VoY3x-NX8dbzfH7unkJ9Lb0@hansken.db.elephantsql.com:5432/oeenwpjw")
	if err != nil {
		log.Fatal("Connect to database was error: ", err)
	}
}

func (d *DB) abort() {
	var err error
	err = d.db.Close()
	if err != nil {
		log.Fatal("Abort from database was error: ", err)
	}
}

func init() {

	Database.Connect()
	fmt.Println("Connecting to database was successful")
}

func main() {

	defer Database.abort()
	result := gin.Default()

	result.GET("/todos", QueryAllHandler)
	result.GET("/todos/:id", QueryByIdHandler)
	result.POST("/todos", InsertHandler)
	result.PUT("/todos/:id", UpdateHandler)
	result.DELETE("/todos/:id", DeleteHandler)

	result.Run(":1234")
}

func QueryAllHandler(c *gin.Context) {

	stmt, err := Database.db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("Cannot prepare query all row statement: ", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Cannot query all todos: ", err)
	}

	var Todos = []Todo{}
	for rows.Next() {

		var t Todo
		var id int
		var title string
		var status string

		err = rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("Cannot scan row into variables: ", err)
		}

		t = Todo{id, title, status}
		Todos = append(Todos, t)
	}

	c.JSON(http.StatusOK, Todos)
}

func QueryByIdHandler(c *gin.Context) {

	stmt, err := Database.db.Prepare("SELECT id, title, status FROM todos WHERE id=$1")
	if err != nil {
		log.Fatal("Cannot prepare query one row statement: ", err)
	}

	var rowId int
	rowId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	row := stmt.QueryRow(rowId)

	var id int
	var title string
	var status string

	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("Cannot scan row into variables: ", err)
	}

	c.JSON(http.StatusOK, Todo{ID: id, Title: title, Status: status})
}

func InsertHandler(c *gin.Context) {

	T := Todo{}
	if err := c.ShouldBindJSON(&T); err != nil {
		log.Fatal("Cannot convert string input to integer: ", err)
	}

	row := Database.db.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", T.Title, T.Status)
	err := row.Scan(&T.ID)
	if err != nil {
		log.Fatal("Cannot scan id: ", err)
	}

	c.JSON(http.StatusCreated, Todo{ID: T.ID, Title: T.Title, Status: T.Status})
}

func DeleteHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Cannot convert string input to integer: ", err)
	}

	stmt, err := Database.db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		log.Fatal("Cannot prepare update statement: ", err)
	}

	if _, err = stmt.Exec(id, "deleted"); err != nil {
		log.Fatal("error excute update: ", err)
	}

	c.JSON(http.StatusCreated, "{\"Status\": \"deleted\"}")
}

func UpdateHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Cannot convert string input to integer: ", err)
	}

	stmt, err := Database.db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		log.Fatal("Cannot prepare update statement: ", err)
	}

	if _, err = stmt.Exec(id, "completed"); err != nil {
		log.Fatal("error excute update: ", err)
	}

	c.JSON(http.StatusCreated, "{\"Status\": \"SUCCESS\"}")
}
