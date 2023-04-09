package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go-lang")

	myCh := make(chan int, 3)
	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Anonymous function 0")
		// fmt.Println(<-myCh, "-0")
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Anonymous function 1")
		myCh <- 10
		myCh <- 20
		myCh <- 600
		// fmt.Println(<-myCh, " - 1")
		// myCh <- 300

		fmt.Println("Done with - 1")
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Anonymous function 2")

		// fmt.Println("fed add values to ch")
		fmt.Println(<-ch, " - 2")
		fmt.Println(<-ch, " - 2")
		fmt.Println(<-ch, " - 2")
		fmt.Println("listened to -2")
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
