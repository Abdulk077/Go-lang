package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "aaple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str1 := "one two three four two two two                 "
	count := strings.Count(str1, "two")
	fmt.Println(count)
	str2 := "    hello world    "
	trimmed := strings.TrimSpace(str2)
	fmt.Println(trimmed)
	// 
	str3 := "Abdul"
	str4 := "Khan"
	result := strings.Join([]string{str3,"wadud",str4}, "")
	fmt.Println(result)
}