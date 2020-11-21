package main

import "fmt"

type param int

const (
	x param = 1
	y param = 2
)

func main() {

	//:= can only be used inside function
	result := Add(x, y)
	fmt.Printf("Sum = %v\n", result)
}

//Declare function
func Add(x param, y param) int {

	return int(x + y)
}
