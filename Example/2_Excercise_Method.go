package main

import "fmt"

type Person struct {
	Name string
}

func (pP *Person) Setter(name string) {
	fmt.Printf("Setter|address P = %p\n", pP)
	pP.Name = name
}

func (pP *Person) Getter() string {
	return pP.Name
}

func (pP *Person) Walk() {
	fmt.Printf("%v walking\n", pP.Name)
}

func (pP *Person) Eat() {
	fmt.Printf("%v eating\n", pP.Name)
}

func (pP *Person) Greeting() {
	fmt.Printf("Hello %v\n", pP.Name)
}

func main() {
	var p Person
	fmt.Printf("main|address p = %p\n", &p)

	p.Setter("Kaioo")

	fmt.Println("Name = ", p.Getter())
	p.Walk()
	p.Eat()
	p.Greeting()
}
