package main



import (
	"fmt"	
    "bufio"	
	"os")
func main() {
	// scanf only take the input until space
	/*fmt.Println("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello", name)*/
	// using bufio to take the input full line
	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello", name)
}