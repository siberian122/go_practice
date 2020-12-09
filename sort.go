package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []int{1, 5, 3, 2, 4}
	sort.Ints(a1)
	fmt.Println(a1)

	a2 := []int{1, 5, 3, 2, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(a2)))
	fmt.Println(a2)

	s1 := []string{"cc", "bb", "aa", "CC", "BB", "AA"}
	sort.Strings(s1)
	fmt.Println(s1)

	s2 := []string{"cc", "bb", "aa", "CC", "BB", "AA"}
	sort.Sort(sort.Reverse(sort.StringSlice(s2)))
	fmt.Println(s2)

	a3 := [5]int{5, 8, 6, 9, 10}
	a4 := a3[:]
	sort.Ints(a4)
	fmt.Println(a4)
}
