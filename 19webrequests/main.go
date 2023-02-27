package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var URL = "https://facebook.com"

func main() {
	fmt.Println("Webrequests : ")

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Printf("Response type is : %T \n", response)

	dataBytes, err := ioutil.ReadAll(response.Body)

	content := string(dataBytes)

	fmt.Println("Data bytes ", content)
}
