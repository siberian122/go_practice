package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("string sample")
	var s1 string
	s2 := "aaa"
	s3 := `bbb
bbbb
bbb`

	fmt.Printf("%v , %T\n", s1, s1)
	fmt.Printf("%v , %T\n", s2, s2)
	fmt.Printf("%v , %T\n", s3, s3)
	fmt.Println(s2 + " + " + s3) // 文字列は「+」で文字連結可能
	fmt.Println(reflect.TypeOf(s2))
	fmt.Println(reflect.TypeOf(s3))
}
