package main

import (
	"fmt"
	"time"
)

func main() {
	currenttime := time.Now()
	fmt.Println(currenttime)
	// 2006/01/02 15:04:05 3:04 AM PM Monday
	// staring date and time
	formated := currenttime.Format("2006-01-02, Monday")
	fmt.Println(formated)
	// convert strint to date
	layout_str := "2006/01/02"
	datestr :="2025-01-25"
	formatted_time, _ := time.Parse(layout_str, datestr)
	fmt.Println(formatted_time)

}