package main

import (
	"encoding/json"
	"fmt"
)



type Person struct {
	Name    string 
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}
func main() {
	fmt.Println(" Json")
	person := Person{ Name: "Abdul", Age:23, IsAdult:true}
	fmt.Println(person)
	// convert person to json encoding
	jsonData, err := json.Marshal(person)
	if err != nil {
		return
	}
	fmt.Println(string(jsonData))
}

