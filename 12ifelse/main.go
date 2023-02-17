package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	loginCount := 10

	if loginCount > 10 {
		fmt.Println("login count is > 10")
	} else if loginCount < 10 {
		fmt.Println("login count is < 10")
	} else {
		fmt.Println("login count is == 10")
	}

	if num := 3; num > 2 {
		fmt.Println("num is greater than 2")
	}
}
