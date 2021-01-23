package main

import (
	"fmt"
	"strconv"
	"strings"
)

const Pi = 3.14

const (
	User = "user"
	pass = "fizzbazz"
)

func main() {
	fmt.Println(Pi, User, pass)
	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "H", "X"))

	var num string = "11"
	i, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(i)
	}

}
