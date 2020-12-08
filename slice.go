package main

import (
	"fmt"
)

func main() {
	var a []int
	fmt.Printf("%v, %T\n", a, a)

	b := []string{"hoge", "fuga"}
	fmt.Println(b[0])

	c := []int{1, 2, 3, 4, 5}
	fmt.Println(len(c))

	d := make([]int, 5)
	//makeでスライスの長さを指定して定義可能
	fmt.Println(d)

	fmt.Println(append(d, 6))

	fmt.Println(d)
	//複数追加
	e := append(d, 6, 7, 8)
	fmt.Println(e)

	var f [][]int
	g := []int{1, 2}
	h := []int{3, 4, 5}
	i := append(f, g, h)
	fmt.Println(append(i, []int{6, 7, 8, 9}))

	//比較はnil年かできない
	x1 := []int{1, 2}
	var x2 []int
	fmt.Println("x1 == nil :", x1 == nil)
	fmt.Println("x2 == nil :", x2 == nil)

	hoge([]int{1, 2})
	hoge([]int{3, 4, 5})
}

func hoge(a []int) {
	fmt.Println(a[0])
}
