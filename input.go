package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//フォーマットを指定し取得
	//%dで整数, %sで文字列を取得

	var a int
	var b string
	fmt.Scanf("%d:%s", &a, &b)
	fmt.Println(a, b)

	var d, e int
	fmt.Scan(&d, &e)
	fmt.Println(d, e)

	//一行取得
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := stdin.Text()
	fmt.Println(s)
}
