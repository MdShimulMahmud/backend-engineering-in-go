package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	var a, b int

	fmt.Println("Enter two numbers: ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	sum := add(a, b)

	fmt.Println("Sum of", a, "and", b, "is", sum)
}
