package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func main() {
	f := pow(3, 2, 10)
	g := pow(3, 3, 20)

	fmt.Println(f, g)
}
