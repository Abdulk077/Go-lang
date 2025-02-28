package main

import "fmt"

func main() {
	day := 3
	// Switch  case for a week
	switch day {

	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Tuesday")
	case 6:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Sunday")

	}
	temperature := 10
	switch {
	case temperature  > 0:
		fmt.Println("liquid")
	case temperature < 0:
		fmt.Println("Frize")
	default:
		fmt.Println("Don't know")
	}
}