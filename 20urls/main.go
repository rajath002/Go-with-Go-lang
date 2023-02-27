package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Welcome to handling URLs in golang")

	const myUrl = "https://linkedin.com:3000/profile?name=rajath&lastname=acharya"

	//parsing
	result, _ := url.Parse(myUrl)

	// fmt.Printf("Type - %T \n", result)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qParams := result.Query()

	fmt.Printf("Data type of qParams: %T \n", qParams)

	for _, v := range qParams {
		fmt.Println(v)
	}

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=rajath",
	}

	anotherUrl := partsOfURL.String()

	fmt.Println(anotherUrl)

}
