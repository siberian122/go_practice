package main

import (
	"fmt"
)

func main() {
	i := 5
	if i < 5 {
		fmt.Println("under 5")
	} else if i == 5 {
		fmt.Println("just 5")
	} else {
		fmt.Println("up 5")
	}
}
