package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Enter the numbers for adding:")
	fmt.Print("Number 1: ")
	fmt.Scan(&a)

	fmt.Print("Number 2: ")
	fmt.Scan(&b)

	c := a + b
	fmt.Println("Result:", c)
}
