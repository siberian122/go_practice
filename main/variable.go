package main

import "fmt"

var ii int = 10

func main() {
	//var 関数外でも宣言可能
	var (
		i    int
		f64  float64
		s    string
		t, f bool = true, false
	)
	fmt.Println(i, f64, t, f)
	// without var
	x := 1
	xs := "string"
	fmt.Println(x, xs)
}
