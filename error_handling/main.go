package main

import "fmt"

func main() {
	// Create a new map
	ans , _ := divide(6, 1)
	fmt.Println(ans)
}

func divide(a, b float64) (float64, error) {
	if b==0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	return a/b, nil
}