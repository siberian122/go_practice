package main

import(
	"fmt"
	"math"
)
func main() {
	var n int
	fmt.Scan(&n)
	var a [n]int
	fmt.Scan(&a)
	var ans int
	for i:=0; i <n-1;i++{
		for l:=i; l<n;l++{
			ans+=math.Abs(int(a[i])-int(a[l]))
	
		}
	}
	fmt.Println(ans)

}
