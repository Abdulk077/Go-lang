package main

import (
	"fmt"
	"os"
)

func main() {
	res, err := http.get(url)
	if err != nil {
		fmt.Println("erros",err)
		return err
	}
	defer res.Body.Close()
	fmt.Printf("Types od response %T\n",res)
	data , err := os.ReadFile(res.Body)
	if err != nil {
		fmt.Println(err)
		return
		
	}
	fmt.Println(string(data))
}