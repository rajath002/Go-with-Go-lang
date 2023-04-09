package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	fmt.Println("Go Routines : ")
	os.Setenv("CGO_ENABLED", "1")
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var arr []int64

	wg.Add(3)
	go func(w *sync.WaitGroup, m *sync.RWMutex) {
		m.Lock()
		arr = append(arr, 1)
		m.Unlock()
		defer w.Done()
	}(wg, mut)

	go func(w *sync.WaitGroup, m *sync.RWMutex) {
		time.Sleep(2 * time.Second)
		m.Lock()
		arr = append(arr, 2)
		m.Unlock()
		defer w.Done()
	}(wg, mut)

	go func(w *sync.WaitGroup, m *sync.RWMutex) {
		m.Lock()
		arr = append(arr, 3)
		m.Unlock()
		defer w.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println("array values ", arr)
}
