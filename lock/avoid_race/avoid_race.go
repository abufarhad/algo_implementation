package main

import (
	"fmt"
	"sync"
)

var count = 0
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	count++
	mutex.Unlock()

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	expectedCount := 1000

	for i := 0; i < expectedCount; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()

	fmt.Println("Expected count - ", expectedCount)
	fmt.Println("Actual count - ", count)

	if expectedCount != count {
		fmt.Println("Race condition detected!")
	} else {
		fmt.Println("No race condition detected.")
	}
}
