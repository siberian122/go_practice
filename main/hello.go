package main

import "fmt"

func init() {
	//mainより先に実行される
	fmt.Println("init")
}

func bazz() {
	fmt.Println("Bazz")
}

func add(x int, y int) int {
	return x + y

}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

//クロージャ―
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func foo(params ...int) {
	fmt.Println(len(params), params)

}
func main() {
	bazz()
	fmt.Println("Hello World!")
	//map (辞書？)
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])

	// 要素の有無の確認
	v, ok := m["apple"]
	fmt.Println(v, ok) // ->100 true

	b := []byte{65, 66}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("AB")
	fmt.Println(c)
	fmt.Println(string(c))

	r := add(10, 20)
	fmt.Println(r)

	f := func() {
		fmt.Println("inner func")
	}

	f()

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	counter := incrementGenerator()
	fmt.Println(counter)
	fmt.Println(counter)

	c1 := circleArea(3.14)
	fmt.Println(c1(2))

	c2 := circleArea(3)
	fmt.Println(c2(2))

	foo(10, 20)
}
