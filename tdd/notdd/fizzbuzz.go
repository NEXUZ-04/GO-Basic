package tdd

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) string {

	fmt.Println("Input = ", n)
	if n%15 == 0 {
		return "FizzBuzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}

func main() {

	fmt.Println(fizzbuzz(1))
	fmt.Println(fizzbuzz(2))
	fmt.Println(fizzbuzz(3))
	fmt.Println(fizzbuzz(4))
	fmt.Println(fizzbuzz(5))
	fmt.Println(fizzbuzz(6))
	fmt.Println(fizzbuzz(7))
	fmt.Println(fizzbuzz(8))
	fmt.Println(fizzbuzz(9))
	fmt.Println(fizzbuzz(10))
	fmt.Println(fizzbuzz(11))
	fmt.Println(fizzbuzz(12))
	fmt.Println(fizzbuzz(13))
	fmt.Println(fizzbuzz(14))
	fmt.Println(fizzbuzz(15))
}
