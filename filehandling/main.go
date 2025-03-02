package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// creating a file using os
	//file, err := os.Create("examples.txt")
	/*if err != nil {
		fmt.Println("Error While Creating File: ", err)
		return
	}
	// closing a file at the end of program
	defer file.Close()
	content := "Hello World By Abdul Wadud Khan"
	// Writing to a file and showing te no of byte
	byte, errors := io.WriteString(file,content)
	if errors != nil{
		fmt.Println("error", errors)
		return
	}
	*/
	// opening a file
	file, err := os.Open("examples.txt")
	if err != nil {
		fmt.Println("Error While Creating File: ", err)
		return
	}
	defer file.Close()
	// to read data we need to use buffer
	buffer := make([]byte,1024)
	for {
		n, err := file.Read(buffer)
		// entering the  end of file
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error While Reading File: ",err)
			return
		}
		fmt.Println(string(buffer[:n]))

	}
	
	fmt.Println("Successfully Read the file")
	//fmt.Println(byte)
	// using os
	content , err := os.ReadFile("examples.txt")
	if err != nil {
		fmt.Println("Error reading", err)
		return 
	}
	fmt.Println(string(content))
}