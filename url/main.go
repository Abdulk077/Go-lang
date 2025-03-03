package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println()
	google := "https://google.com/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of url%T", google)

	personalURL, err := url.Parse(google)
	if err != nil {
		fmt.Println("cant parse url")
		return
	}
	fmt.Println(personalURL)
	// printing parsed url
	fmt.Println(personalURL.Scheme)
	fmt.Println(personalURL.Host)
	fmt.Println(personalURL.Path)
	fmt.Println(personalURL.RawQuery)
	// modifying url 
	personalURL.Path = "/newPath"
	personalURL.RawQuery = "username=iamabdul"
	// Constructing a new url string
	newurl := personalURL.String()
	fmt.Println(newurl)

}