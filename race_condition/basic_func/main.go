package main

import (
	"fmt"
	"sync"
)

var (
	cnt   int
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func increment() {
	mutex.Lock() //to avoid condition
	cnt++
	mutex.Unlock()
	wg.Done()
}

func main() {
	numOfRoutine := 5
	wg.Add(numOfRoutine)
	for i := 0; i < numOfRoutine; i++ {
		go increment()
	}

	wg.Wait()
	fmt.Println("counter ", cnt)
}
