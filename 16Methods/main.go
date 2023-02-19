package main

import "fmt"

func main() {

	println("All about Structs")

	user := User{"Rajath", "rj@go.dev", "28", true}

	fmt.Println("User : ", user)

	fmt.Printf("user details %+v \n", user)

	fmt.Printf("User name is %v and email is %v\n", user.Name, user.Email)

	user.printName()
	user.printUserStatus()
	user.changeUserEmail()

	fmt.Println("changed email ", user.Email)

}

type User struct {
	Name   string
	Email  string
	Age    string
	Status bool
}

func (u User) printName() {
	fmt.Println("User name is : ", u.Name)
}

func (u User) printUserStatus() {
	fmt.Println("User status is : ", u.Status)
}

func (u User) changeUserEmail() {
	u.Email = "test@go.dev"
	fmt.Println("changed user email : ", u.Email)
}
