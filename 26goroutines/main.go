package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually these are pointers

func main() {
	// go greeter("Hello")
	// greeter("world")

	websiteslist := []string{
		"https://google.com",
		"https://fb.com",
		"https://instagram.com",
	}

	for _, endpoint := range websiteslist {
		go getStatusCode(endpoint)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(i, s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.StatusCode, endpoint)
}
