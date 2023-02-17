package main

import "fmt"

func main() {
	var fruits [5]string

	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[2] = "Orange"
	fruits[3] = "Quava"
	fruits[4] = "Papaya"

	fmt.Println("Fruits in the list : ", fruits, " Length is ", len(fruits))

	var animals = []string{"Lion", "Tiger"}

	fmt.Println("Animals in the list", animals, ", Length is ", len(animals))

}
