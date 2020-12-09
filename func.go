package main

import (
	"fmt"
	"strings"
)

func hoge() {
	//処理
	fmt.Println("hoge")

}
func Hoge() {
	fmt.Println("Hoge")
}

func hoge1(num int) {
	fmt.Println(num)
}

//型が違う場合
func hoge2(num1 int, num2 float32) {
	fmt.Println(float32(num1) + num2)
}

//型が一緒
func fuga(num1, num2 int) {
	fmt.Println(num1 - num2)
}

//返却値がある場合
func hoge3(num int) string {
	str := strings.Repeat("Go", num)
	return str
}

//複数返却
func hoge4(num1, num2 int) (string, string) {
	str1 := strings.Repeat("Go", num1)
	str2 := strings.Repeat("GO", num2)
	return str1, str2
}

//引数を...方にすると引数を可変に
func hoge5(s ...string) {
	for i, v := range s {
		fmt.Println(i, v)
	}
}

//使われていない変数があるとエラーに

func main() {
	hoge()
	Hoge()
	hoge1(10)
	hoge2(2, 3.5)
	fuga(10, 5)
	str := hoge3(3)
	fmt.Println(str)
	str1, str2 := hoge4(5, 10)
	fmt.Println(str1)
	fmt.Println(str2)
	hoge5("aa")
	hoge5("bb", "cc", "dd")
}
