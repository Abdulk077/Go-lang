package main 

import 
	"fmt"

func main() {
	fmt.Println("Hello, World!")
	// for variables
	var name  string = "Abdul"
	fmt.Println(name)
	// for float
	var dimension float64 = 67.44
	fmt.Println(dimension)
	// for boolean
	var isTrue bool = true
	fmt.Println(isTrue)
	// for a constant variable
	const pi float64 = 3.14159	
	fmt.Println(pi)
	// using short hand for variables
	age := 23
	fmt.Println(age)
	// Examples of public and private variables
	// Public variables
	var PublicVariable string = "I am a public variable"
	fmt.Println(PublicVariable)
	// Private variables
	var private string = "I am a private variable"
	fmt.Println(private)		

}