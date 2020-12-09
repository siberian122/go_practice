package main

import (
	"fmt"
)

func main() {
	// スライスをスライス
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Printf("ex.1 tyoe: %T, value: %v\n", s[0:], s[0:])
	fmt.Printf("ex.2 type: %T, value: %v\n", s[:3], s[:3])
	fmt.Printf("ex.3 type: %T, value: %v\n\n", s[1:3], s[1:3])

	// 配列スライス
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("ex/4 type: %T, value: %v\n\n", a[1:3], a[1:3])

	//文字列スライス
	str := "aiueo"
	fmt.Printf("ex.5 type: %T, value: %v\n", str[3:], str[3:])
}
