package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter a number to check even or odd:")
	fmt.Scan(&a)

	if a%2 == 0 {
		fmt.Println("Number", a, "is an even number")
	} else {
		fmt.Println("Number", a, "is an odd number")
	}
}
