package main

import "fmt"

func main() {

	var s = []int{0, 1}
	//Or use "make" for slice without initial value
	//s := make([]int, 5)
	fmt.Printf("%T => %v\n", s, s)

	//Expand size of slice with "append"
	//!Cannot do like this --> s[6] = 6
	s = append(s, 5)
	fmt.Printf("%T => %v\n", s, s)

	//length of slice = 3 (0, 1, 5)
	l := len(s)
	fmt.Println(l)

	//capability of slice = 4 because s (0, 1, 5, [appended position])

	//***slice process = create array before insert value
	c := cap(s)
	fmt.Println(c)

	a := []int{0, 1, 2, 3, 4, 5}
	a = a[1:4] //a = [1,2,3] but capability count from start till end of original slice/array [1,2,3,_,_] (Cap = 5)
	fmt.Printf("value = %v, len = %d, cap = %d\n", a, len(a), cap(a))

	a = a[:2] //a = [1,2] but capability count from start till end of original slice/array [1,2,_,_,_] (Cap = 5)
	fmt.Printf("value = %v, len = %d, cap = %d\n", a, len(a), cap(a))

	a = a[1:] //a = [2] but capability count from start till end of original slice/array [2,_,_,_,_] (Cap = 4)
	fmt.Printf("value = %v, len = %d, cap = %d\n", a, len(a), cap(a))
}
