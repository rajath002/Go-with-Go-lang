package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("More on JSON")

	// EncodeJsonData()
	DecodeJsonData()

	// People()
}

type Course struct {
	Name     string `json:"courseName"`
	Platform string `json:"website"`
	Rating   int
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJsonData() {
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

func DecodeJsonData() {
	jsonDataFromWeb := []byte(`{
		"courseName": "ReactJS",
		"website": "Nodejs, React",
		"Rating": 5,
		"tags": ["Reactjs","JS"]
	}`)
	var lcoCourse Course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		err := json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		if err != nil {
			fmt.Print(err)
			panic(err)
		}
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("invalid json data", checkValid)
	}

	// Some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}

	err := json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	fmt.Printf("--> %#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T\n", k, v, v)
	}
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func People() {
	jsonStr := `[{"name":"Alice","age":30},{"name":"Bob","age":25}]`

	var people []Person

	err := json.Unmarshal([]byte(jsonStr), &people)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for _, person := range people {
		fmt.Printf("%s is %d years old.\n", person.Name, person.Age)
	}
}
