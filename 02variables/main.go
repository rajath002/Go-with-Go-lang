package main

import "fmt"

const LoginToken string = "logingtoken" // Public

func main() {
	username := "Rajath"
	fmt.Println("hello", username)
	fmt.Printf("Varibles of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println("Is user logged in : ", isLoggedIn)
	fmt.Printf("Varibles of type : %T \n", isLoggedIn)

	var smallFloat float64 = 23.4532223445465465
	fmt.Println("Float 64: ", smallFloat)
	fmt.Printf("Varibles of type : %T \n", smallFloat)

	// Default values and some aliases
	var anotherVar int
	fmt.Println("anotherVar int : ", anotherVar)
	fmt.Printf("Varibles of type : %T \n", anotherVar)

	//implicit type
	var website = "rajath.com"
	fmt.Println("website string : ", website)
	fmt.Printf("Varibles of type : %T \n", website)

	fmt.Println("Login token", LoginToken)
	fmt.Printf("Varibles of type : %T \n", LoginToken)
}
