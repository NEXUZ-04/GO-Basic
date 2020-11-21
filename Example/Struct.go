package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	//Declare instant v
	v := Vertex{X: 2, Y: 1}

	fmt.Printf("%#v\n", v)
	fmt.Println(v.X)

	//Declare pointer
	var pv *Vertex
	pv = &v

	(*pv).X = 39
	fmt.Println(v.X)
}
