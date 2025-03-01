package main

import "fmt"
func add (a,b int) int{
	return a + b
}
func main() {
	fmt.Println("hii")
	//data := add(4,6)
	defer fmt.Println("hello world", )
	defer fmt.Println("Middle")
	fmt.Println("Welcome")
}