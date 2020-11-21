package main

import "fmt"

func main() {

	type student struct {
		Name  string
		Class string
	}

	m := map[int]student{
		1: student{Name: "A", Class: "Engineer"},
		2: student{Name: "B", Class: "Programer"}, //Don't forget last comma
	}

	fmt.Println(m)
}
