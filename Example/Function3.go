package main

import (
	"fmt"
	"math"
)

func main() {
	plus := func(x float64, y float64) float64 {
		return x + y
	}

	a := compute(plus)
	fmt.Println(a)

	r := compute(math.Pow)
	fmt.Println(r)

	x := compute(multiply)
	fmt.Println(x)
}

func multiply(x float64, y float64) float64 {
	return x * y
}

func compute(fn func(float64, float64) float64) float64 {
	r := fn(3, 4)
	return r
}
