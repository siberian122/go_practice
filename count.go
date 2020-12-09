package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaabbbcccaaabbbccc"
	s1 := strings.Count(s, "a")
	fmt.Println(s1)
}
