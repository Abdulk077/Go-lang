package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func forgetrequest (){
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
func forPostrequest() {
	todo := Todo {
		UserID: 23,
		Title: "Abdul",
		Completed: true,
	
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		return 
	}
	// convert jason data to string
	jsonString := string(jsonData)
	jsonReader, err := string.NewReader(jsonString)
	myurl := "http://jasonplaceholder.typicode.com/todos"
	res, err := http.Post(myurl, "application/json", jsonReader)
	if err != nil {
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string( data ))

}
type Todo struct {
	UserID int `json:"user_id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}
func main() {

	forPostrequest()
}