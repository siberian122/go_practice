package main

import (
	"fmt"
)

//個別宣言
var a int
var b int = 1
var c = 2

//まとめて宣言
var (
	d int
	e int = 3
	f     = 4
)
var a1, a2 int
var b1, b2 int = 5, 6
var c1, c2 = 7, 8

func main() {
	aa := 9
	bb, cc := 10, 11
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(a1, a2)
	fmt.Println(b1, b2)
	fmt.Println(c1, c2)
	fmt.Println(aa)
	fmt.Println(bb, cc)
}
