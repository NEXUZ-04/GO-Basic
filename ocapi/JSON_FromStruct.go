package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	ID     int    `json:"MyId"`
	Title  string `json:"MyTitle"`
	Status string `json:"MyStatus"`
}

func main() {

	t := Todo{
		ID:     1,
		Title:  "Demon Slayer",
		Status: "Sold out",
	}

	b, err := json.Marshal(t)
	fmt.Printf("Type = %T\n", b)
	fmt.Printf("Value = %v\n", b)
	fmt.Printf("String value = %s\n", b)
	fmt.Println(err)
}
