package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("number samlpe")
	var a int
	var b float64
	var c float64 = 1.0
	var d float64 = 1.0

	fmt.Println(a, b)
	fmt.Println(c + d)

	fmt.Println(float64(a) + c + d) //型をそろえる
	fmt.Println(reflect.TypeOf(a))  //reflect.TypeOfで型確認
	fmt.Printf("%v , %T\n", a, a)   //Printfで％Tを使うと型出力
	fmt.Println()

	var e int
	e = 10
	fmt.Println(e)
	fmt.Println(e - 11)

	var f uint
	f = 10
	fmt.Println(f)
	fmt.Println(f - 11)

}
