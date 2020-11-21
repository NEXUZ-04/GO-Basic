package main

import (
	"fmt"
	"runtime"
	"time"
)

//Switch case : No need break!
//              fallthrough = go to expression of the next case

func main() {

	//Normal switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "Linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	//Switch with condition
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Far away.")
	}
}
