package main

import "fmt"

func main() {

	a, b := Swap("Hello", "World")
	fmt.Println(a, b)
}

//Declare function return muliple result
func Swap(x string, y string) (string, string) {

	return y, x
}
