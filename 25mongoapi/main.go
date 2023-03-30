package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rajath002/gomongoapi/router"
)

func main() {
	fmt.Println("Go mongodb API")

	r := router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}
