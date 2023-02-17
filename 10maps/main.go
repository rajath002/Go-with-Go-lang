package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Maps in Go-Lang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["TS"] = "Typescript"

	fmt.Println("List of all languages", languages)
	fmt.Println("Type key of the language to remove: ")

	// Un comment below code if you want to read the user input
	/*
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error ", err)
		}
	*/
	input := "rb"

	userInput := strings.ToUpper(strings.TrimSpace(input))
	fmt.Println("Removing ", userInput, languages[userInput])

	delete(languages, userInput)

	fmt.Println("After Removing Ruby the  languages remaining is ", languages)

	// Courses are
	for k, v := range languages {
		fmt.Printf("(%v) %v \n", k, v)
	}

}
