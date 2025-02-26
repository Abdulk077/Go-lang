package main

import "fmt"

// Function to calculate the sum of two numbers
func sum(a int, b int) int {
	return a + b
}

// Function to calculate the product of two numbers
func product(a int, b int) int {
	return a * b
}

func main() {
	// calling the sum function
	fmt.Println(sum(10, 20))
	// calling the product function
	fmt.Println(product(10, 20))
}
