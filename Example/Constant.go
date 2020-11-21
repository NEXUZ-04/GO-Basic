package main

import "fmt"

/*
"iota"

^
^
^
^
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Satureday = 5
	Sunday    = 6
)
*/

func main() {

	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Satureday
		Sunday
	)

	fmt.Println("Monday: ", Monday)
	fmt.Println("Tuesday: ", Tuesday)
	fmt.Println("Wednesday: ", Wednesday)
	fmt.Println("Thursday: ", Thursday)
	fmt.Println("Friday: ", Friday)
	fmt.Println("Satureday: ", Satureday)
	fmt.Println("Sunday: ", Sunday)
}
