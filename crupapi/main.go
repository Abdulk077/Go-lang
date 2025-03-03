package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Using Get Method
	fmt.Println()
	res, err := http.Get("url")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("error: ", res.Status)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(data)

	
}