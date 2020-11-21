package main

import "fmt"

func main() {
	i := 5
	f := float64(i)
	u := uint(f)

	fmt.Println(i, f, u)
	fmt.Printf("%T %T %T\n", i, f, u)
}
