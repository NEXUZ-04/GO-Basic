package main

import "fmt"

type Dog struct {
	name string
}

type Cat struct {
	name string
}

type Waker interface {

	//Define behavior
	Walk()
}

func main() {

	var param interface{}

	param = "Kaioo"
	Display(param) //Value: Kaioo Type: string

	param = 5
	//Cannot do like : param = param + 1 cuz param isn't an integer
	Display(param) //Value: 5 Type: int (interface just shows a  type of what it collects)

	iparam := param.(int) //iparam will be actual int
	iparam = iparam + 1
	Display(iparam) //Value: 6 Type: int

	D := Dog{name: "Luffy"}
	Walk(D)

	C := Cat{name: "Zoro"}
	Walk(C)
}

func Display(input interface{}) {

	fmt.Printf("Value: %#v, Type: %T\n", input, input)
}

func (D Dog) Walk() {

	fmt.Printf("%#v is walking like a dog\n", D.name)
}

func (C Cat) Walk() {

	fmt.Printf("%#v is walking like a cat\n", C.name)
}

//Both Dog and Cat can use Walk()
func Walk(W Waker) {

	W.Walk()
}
