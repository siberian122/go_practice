package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a = make(map[string]int)
	fmt.Println("a : %v , T\n", a, a)
	fmt.Println(reflect.TypeOf(a))

	//マップの定義
	b := map[string]int{"a": 1, "b": 2}
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Println(reflect.TypeOf(b))

	b["b"] = 10
	fmt.Println(b["a"], b["b"])
	//存在しない要素を指定してもエラーにならない
	fmt.Println(b["c"])

}
