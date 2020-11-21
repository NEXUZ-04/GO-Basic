package main

import "fmt"

func main() {
	names := []string{"game", "bank", "samui", "dog", "jiew"}

	for i, name := range names {
		fmt.Println(i, name)
	}
}
