package main

import "fmt"

func main() {
	inc, cur := adder()

	fmt.Println(cur()) //0
	fmt.Println(inc()) //1
	fmt.Println(inc()) //2
	fmt.Println(cur()) //2
	fmt.Println(inc()) //3
	fmt.Println(cur()) //3
	fmt.Println(cur()) //3
	fmt.Println(inc()) //4
}

func adder() (func() int, func() int) {
	sum := 0

	return func() int {
			sum += 1
			return sum
		},
		func() int {
			return sum
		}
}
