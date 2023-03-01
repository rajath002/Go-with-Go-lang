package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Webrequests ")
	const url = "https://dummyjson.com/products/1"
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Printf("response body type : %T\n", response.Body)

	var stringBuilder strings.Builder

	byteCount, _ := stringBuilder.Write(content)

	fmt.Println("Bytecount is ", byteCount)

	fmt.Println(stringBuilder.String())

	// fmt.Println("response", string(content))
}
