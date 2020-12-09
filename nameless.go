package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	func() {
		fmt.Print("nameless")
	}()

	nameless1 := func() {
		fmt.Println("function")
	}
	nameless1()

	nameless2 := func(a, b int) (string, string) {
		aStr := strings.Repeat("GO", a)
		bStr := strings.Repeat("go", b)
		return aStr, bStr
	}
	aa, bb := nameless2(3, 5)
	fmt.Println(aa, bb)

	a := 5
	go func(a int) {
		for i := 0; i < 5; i++ {
			fmt.Println("hoge", i)
			time.Sleep(1 * time.Second)
		}
	}(a)
	time.Sleep(time.Duration(a) * time.Second)
}
