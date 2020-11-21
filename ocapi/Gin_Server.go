package main

import (
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Todos = []Todo{}

//Parameter name must be capital letter
type Todo struct {
	ID int         `json:"id"`
	Title string   `json:"title"`
	Status string  `json:"status"`
}

func getHandler(c *gin.Context) {

	status := c.Query("status")
	items := []Todo{}

	for _, t := range Todos {
		if t.Status == status {
			items = append(items, t)
		}
	}

	c.JSON(http.StatusOK, items)
}

func getByIdHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	for _, t := range Todos {
		if t.ID == id {
			c.JSON(http.StatusOK, t)

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

func postHandler(c *gin.Context) {
	
	T := Todo{}
	if err := c.ShouldBindJSON(&T); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	Todos = append(Todos, T)
	c.JSON(http.StatusCreated, "Todo was created.")
}

func deleteHandler(c *gin.Context) {
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	NewTodos := []Todo{}
	for _, t := range Todos {
		if t.ID != id {
			NewTodos = append(NewTodos, t)
		}
	}

	Todos = NewTodos
	c.JSON(http.StatusCreated, "Todo was deleted.")
}

func main() {
	r := gin.Default()
	r.GET("/todos", getHandler)
	r.GET("/todos/:id", getByIdHandler)
	r.POST("/todos", postHandler)
	r.DELETE("/todos/:id", deleteHandler)
	r.Run(":1234")
}