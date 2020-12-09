package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaabbbcccaaabbbccc"
	s1 := strings.Replace(s, "bbb", "XXX", -1)
	fmt.Printf("%v\n%v\n", s, s1)
}
