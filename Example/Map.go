package main

import "fmt"

func main() {

	//Declare variable
	var m map[string]string

	//Allocate memory
	m = make(map[string]string)

	//Can use only --> m := map[string]string

	//Try print init value
	fmt.Println(m)

	//Create key
	m["Name"] = "Kaioo"
	v := m["Name"]
	fmt.Println(m)
	fmt.Println(v)

	//Delete key
	delete(m, "Name")
	v = m["Name"]
	fmt.Println(m)
	fmt.Println(v) //Empty or 0 if value in key was INT
}
