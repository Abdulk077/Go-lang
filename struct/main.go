package main

import "fmt"

type Person struct {
	Name string
	age int
}
type Address struct {
	Area string
}
type Contact struct {
	Email string
}
type Employee struct {
	Person_details Person
	Person_contact Contact
	Person_address Address
}
func main() {

	person1 := Person{
		Name: "John",
		age: 20,
	}
	fmt.Println(person1)
	fmt.Println(person1.age)

	//

	var employe1  Employee
	employe1.Person_details = Person{
		Name: "Abdul",
		age: 21,
	}
	employe1.Person_contact.Email = "abdul@gmail.com"
	employe1.Person_address.Area = "BBSR"
	fmt.Println(employe1)
	
}