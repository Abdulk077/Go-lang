package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Leaning go routines")
	go sayhello()
	go sayhii()
	time.Sleep(1500 * time.Millisecond)

}
func sayhello() {
	fmt.Println("Hello, world!")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Say Hello function eneded ")
}
func sayhii(){
	fmt.Println("hii Abdul")
	time.Sleep(1600 * time.Millisecond)
	fmt.Println("Say hii func eneded ")

}