package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for out Pizza: ")

	// Comma ok || err ok
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for reading, ", input)
	fmt.Printf("Input type is : %T, ", input)
}
