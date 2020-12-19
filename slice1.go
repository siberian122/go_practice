package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3}
	b := a

	fmt.Printf("a:%v, %v\n", a[0], &a[0])
	fmt.Printf("b:%v, %v\n", b[0], &b[0])

	a[0] = 10
	fmt.Printf("a:%v, %v\n", a[0], &a[0])
	fmt.Printf("b:%v, %v\n", b[0], &b[0])

	 // スライスした場合も同じアドレスになる
    c := a[:]
    fmt.Printf("c:%v, %v\n", c[0], &c[0])
    a[0] = 20
    fmt.Printf("c:%v, %v\n", c[0], &c[0])
    fmt.Println()

    // 一部分をスライスしても同じ
    d := a[:2]
    fmt.Printf("d:%v, %v\n", d[0], &d[0])
    a[0] = 30
    fmt.Printf("d:%v, %v\n", d[0], &d[0])
    fmt.Println()

    // スライスを作成する際に第3引数でスライスの容量を指定できる。
    // 容量を超えるまでは同じアドレスを使用し、超えると新しく作り直される
    g := make([]int, 3, 3)
    fmt.Printf("g:%v, %v, %v\n", g, g[0], &g[0])
    g[0] = 1
    fmt.Printf("g:%v, %v, %v\n", g, g[0], &g[0])
    g1 := append(g, 4)
    fmt.Printf("g1:%v, %v, %v\n", g1, g1[0], &g1[0])
    fmt.Println()
    // 新しいアドレスが割り当てられたことで変更してもコピー元には影響しない
    g1[1] = 10
    fmt.Printf("g:%v, %v, %v\n", g, g[1], &g[0])
    fmt.Printf("g1:%v, %v, %v\n", g1, g1[1], &g1[0])
    fmt.Println()

    // capを多めにとった場合
    h := make([]int, 3, 5)
    fmt.Printf("h:%v, %v, %v\n", h, h[0], &h[0])
    h[0] = 1
    fmt.Printf("h:%v, %v, %v\n", h, h[0], &h[0])
    h = append(h, 4)
    fmt.Printf("h:%v, %v, %v\n", h, h[0], &h[0])
    fmt.Println()

    // 関数に渡す際もメモリ上のアドレスを渡している
    z := cp(a)
    fmt.Println("z:", &z[0])
    fmt.Println("z:", z[0])
}

func cp(s []int) []int {
    fmt.Println("s:", &s[0])
    return s
}