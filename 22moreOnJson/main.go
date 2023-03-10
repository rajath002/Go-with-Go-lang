package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("More on JSON")

	createCourse()
}

type Course struct {
	Name     string `json:"courseName"`
	Platform string `json:"website"`
	Rating   int
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func createCourse() {
	var course = []Course{
		{Name: "ReactJS", Platform: "Nodejs, React", Tags: []string{"Reactjs", "JS"}, Rating: 5, Password: "js@123"},
		{Name: "Python", Platform: "Python", Tags: []string{"Python", "PY"}, Rating: 5, Password: "py@123"},
		{Name: "Go", Platform: "go", Tags: []string{}, Rating: 4, Password: "go@123"},
	}
	fmt.Printf("%T", course)

	data, err := json.MarshalIndent(course, "", "\t")

	if err != nil {
		panic(err)
	}

	jsonData := string(data)
	fmt.Println("The value is", jsonData)
}
