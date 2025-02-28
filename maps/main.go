package main

import "fmt"

func main(){
	fmt.Println(0)

	students := make(map[string]int)

	students["Abdul"] = 100
	students["Khan"] = 50
	students["Chan"] = 89
	students["Sammad"] = 99
	fmt.Println(students["Abdul"])
	students["Khan"] = 100
	fmt.Println(students["Khan"])
	// To Delete Any Key
	delete(students, "Chan")
	marks , exist := students["Abdul"]
	fmt.Println(marks)
	fmt.Println(exist)

}