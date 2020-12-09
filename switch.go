package main

import (
	"fmt"
)

func main() {
	num := 1
	switch num {
	case 1:
		fmt.Println("1!")
		fallthrough
	case 2:
		fmt.Println("2!")
	default:
		fmt.Println("default")
	}
}
