package main

import "fmt"

// Defer

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println("\nOne")

	defer myDefer()

	defer fmt.Println("Two")
	defer fmt.Println("Three")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

/*
**Output**
Hello
Three
Two
43210
One
World

 ===Stack===
 -----------------------------------------
 [ World, One, 0, 1, 2, 3, 4,  Two, Three]
 -----------------------------------------
*/
