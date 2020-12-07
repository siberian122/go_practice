package main

import (
	"fmt"
)

var global int

func main() {
	var local int
	fmt.Println(global)
	fmt.Println(local)
	fmt.Println("hello world")
}
