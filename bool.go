package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("bool sample")
	var a bool
	fmt.Printf("%v , %T\n", a, a)
	fmt.Printf(reflect.TypeOf(a))
}
