package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//数字を数値に変換
	a1 := "1"
	a2, a3 := strconv.Atoi(a1)
	fmt.Println(a1, reflect.TypeOf(a1))
	fmt.Println(a2, reflect.TypeOf(a2), a3)

	a4, a5 := strconv.Atoi("a")
	fmt.Println(a4, a5)

	//数値を数字に変換
	b := 1
	b1 := strconv.Itoa(b)
	fmt.Println(b, reflect.TypeOf(b), b1, reflect.TypeOf(b1))
}
