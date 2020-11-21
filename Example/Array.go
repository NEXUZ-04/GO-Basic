package main

import "fmt"

func main() {
	var a [3]string

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	//Convert Array to Slice (Array without size)
	//primes := [6]int{2, 3, 5, 7, 11, 13} --> Array
	//primes := []int{2, 3, 5, 7, 11, 13, 17} --> Slice

	var s []int = primes[1:4]
	//[1:4] = Select index 1,2,3 in primes[6]
	//[:4]  = Select index 0,1,2,3 in primes[6]
	//[1:]  = Select index 1,2,3,4,5 in primes[6]
	fmt.Printf("%T => %v\n", s, s)

	//actual length of slice
	l := len(s)
	fmt.Println(l)

	//If slice was changed, values of original array and other(s) slice will be also changed
	s[0] = 100
	fmt.Printf("%T => %v\n", s, s)
	fmt.Printf("%T => %v\n", primes, primes)
}
