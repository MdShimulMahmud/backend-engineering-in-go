package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	p := person{
		firstName: "John",
		lastName:  "Doe",
		age:       25,
	}

	fmt.Println("First name:", p.firstName)
	fmt.Println("Last name:", p.lastName)
	fmt.Println("Age:", p.age)

}
