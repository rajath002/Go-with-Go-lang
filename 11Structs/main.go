package main

import "fmt"

func main() {

	println("All about Structs")

	user := User{"Rajath", "rj@go.dev", "28"}

	fmt.Println("User : ", user)

	fmt.Printf("user details %+v \n", user)

	fmt.Printf("User name is %v and email is %v", user.Name, user.Email)
}

type User struct {
	Name  string
	Email string
	Age   string
}
