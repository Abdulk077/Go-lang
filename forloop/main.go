package main

import "fmt"

func main() {
	fmt.Println()
	// for Loop 
	for i := 0; i < 5 ; i++ {
		fmt.Println(i)
	}

	// Using Range
	number :=[5]int{1, 2, 3, 4,5}
	for index , value := range number{
		fmt.Println(index)
		fmt.Println(value)
	}
	// with String
	data := "HELLO"
	for index, value := range data {
		fmt.Println(index)
		fmt.Println(value)
	}
}