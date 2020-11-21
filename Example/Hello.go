package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	//defer = execute when finish function
	defer fmt.Println("d")
	defer fmt.Println("l")
	defer fmt.Println("r")
	defer fmt.Println("o")
	defer fmt.Println("w")

	fmt.Println("hello")
}
