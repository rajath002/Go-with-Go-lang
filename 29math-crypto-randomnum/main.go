package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	/*
		var myNumberOne int = 2
		var myNumberTwo float64 = 4.5

		fmt.Println("The sum is : ", myNumberOne+int(myNumberTwo))
	*/

	/*
		//random number
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Intn(5))
		fmt.Println(rand.Intn(5))
		fmt.Println(rand.Intn(5))
	*/

	// Random with Crypto Package
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(9))

	fmt.Println("random int ", myRandomNum)

}
