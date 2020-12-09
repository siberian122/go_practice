package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	for i, v := range []int{5, 6, 7} {
		fmt.Printf("i =%d, v =%d \n", i, v)
	}
	fmt.Println()

	//while文的に
	w := 0
	for w < 2 {
		fmt.Println(w)
		w++
	}
	fmt.Println()

	//無限ループ
	a := 0
	for {
		a++
		if a > 5 {
			break
		} else if (a % 2) == 0 {
			continue
		}
		fmt.Println(a)
	}
}
