package main

import (
	"fmt"
	"sync"
)

var mtx sync.Mutex

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	func() {
		defer wg.Done()
		n := 2
		fmt.Println("Goroutine 1 ")

		mtx.Lock()
		defer mtx.Unlock()
		if n%2 == 0 {
			return
		}
	}()

	func() {
		defer wg.Done()
		n := 2
		fmt.Println("Goroutine 2 ")

		mtx.Lock()
		defer mtx.Unlock()
		if n%2 == 0 {
			return
		}
	}()

	wg.Wait()
	fmt.Println("Main routine complete")
}
