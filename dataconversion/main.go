package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println()
	// int to float
	var num int = 40
	var data float64 = float64(num)
	fmt.Println(data)
	// int to string value
	str := strconv.Itoa(num)
	fmt.Println(str)
	// 

}