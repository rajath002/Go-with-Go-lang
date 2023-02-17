package main

import "fmt"

func main() {
	varNumber := 10

	ptr := &varNumber

	fmt.Println("Pointer address", ptr)

	fmt.Println("Pointer value ", *ptr)

	var ptr2 *string

	fmt.Println("Pointer address", ptr2)

	str := "hello world"

	ptr2 = &str

	fmt.Print("Pointer value ", *ptr2)
}
