package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaa,bbb,ccc,aaa,bbb,ccc"
	s1 := strings.Split(s, ",")
	fmt.Printf("%v\n%v\n", s, s1)
}
