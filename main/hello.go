package main

import "fmt"

func init() {
	//mainより先に実行される
	fmt.Println("init")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	bazz()
	fmt.Println("Hello World!")
}
