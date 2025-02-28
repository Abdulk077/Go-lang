package main

import "fmt"

func main(){
	numbers := []int{1, 2, 3, 4, 5}	
	// apending numbers
	numbers = append(numbers, 6,4,7,2,9,55,3)
	// useing make 
	number := make([]int,2,5)
	number = append(number,5)
	number = append(number,4)
	number = append(number,4)


	fmt.Println("cap" , cap(number))
	fmt.Println(len(number))
	fmt.Println(number)
}