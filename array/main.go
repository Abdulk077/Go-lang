package main

import "fmt"

func main() {
	// Creating an array
	var names [3]string
	names[0] = "Abdul"
	names[1] = "Wadud"
	names[2] = "Khan"
	fmt.Println(names)
	// enter input during initialization
	var ages = [3]int{23, 24, 25}
	fmt.Println(ages)

	// print by index
	fmt.Println(names[0])
}