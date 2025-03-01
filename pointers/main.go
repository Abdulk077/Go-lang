package main

import "fmt"

func modifiaction(num *int){
	*num = *num * 3
}
func main() {
	var num int

	ptr := &num
	fmt.Println(*ptr)
	fmt.Println(ptr)
	value := 20
	fmt.Println(value)
	modifiaction(&value)
	fmt.Println(value)
}