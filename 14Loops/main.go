package main

import "fmt"

func main() {
	languages := []string{"Java", "Javascript", "Typescript", "Golang"}

	// for i := 0; i < len(languages); i++ {
	// 	fmt.Println("Language ", i, languages[i])
	// }

	// for i, v := range languages {
	// 	fmt.Printf("key %v and value is %v \n", i, v)
	// }

	for i := 0; i < len(languages); i++ {
		if i == 3 {
			fmt.Println("Breaking...")
			break
		}
		if i == 2 {
			// fmt.Println("continue...")
			// continue
			fmt.Println("goto label...")
			goto label
		}
		fmt.Println("language is", languages[i])
	}

	fmt.Println("Loop is over")

label:

	fmt.Println("this is from Rajath, thank you")
}
