package main

import (
	"fmt"
)

func main() {
	//配列定義
	var a [3]int
	//定義のみ➞ゼロ値が設定
	fmt.Printf("%v , $T\n", a, a)

	b := [2]string{"aaa", "bbb"}
	fmt.Println(b[1])

	//初期値を指定すると「...」で長さを省略できる
	c := [...]int{1, 2}
	fmt.Println(len(c))

	//インデックスを指定して値を代入
	d := [3]int{1: 1, 2: 10}
	fmt.Println(d)

	//2重配列
	var e [3][2]int
	fmt.Printf("%v ,%T", e, e)
	e[1][0] = 2
	e[1][1] = 3
	fmt.Println(e)

	f := [3][2]int{{0, 1}, {2, 3}, {4, 5}}
	fmt.Println(f)

	x1 := [2]int{1, 2}
	x2 := [2]int{1, 2}
	fmt.Println("x1 == x2", x1 == x2)
	x2[0] = 3
	fmt.Println("x1 ==x2", x1 == x2)

	hoge([2]int{1, 2})

}
func hoge(a [2]int) {
	fmt.Println(a[0])
}
