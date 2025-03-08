package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Sync in Go")
	var wg sync.WaitGroup
	for i:= 1;  i <= 3; i++{
		wg.Add(1)
		go worker(i,&wg)
	}
	wg.Wait()
	fmt.Println("Task complete")

}
func worker(i int , wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d finished\n", i)
}