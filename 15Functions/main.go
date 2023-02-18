package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Functions")
	greetings()

	fmt.Println("Sum is ", sum(5, 10))

	result, message := dynamicSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(message, result)
}

func greetings() {
	fmt.Println("Greetings")
}

func sum(a int, b int) int {
	return a + b
}

func dynamicSum(args ...int) (int, string) {
	result := 0

	for _, v := range args {
		result += v
	}

	return result, "Total Sum is :"
}
