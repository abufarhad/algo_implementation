package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sharedInt   int = 0
	unusedValue int = 0
	once        sync.Once
	mutex       sync.Mutex
)

func runSimpleReader() {
	mutex.Lock()
	defer mutex.Unlock()

	once.Do(func() {
		val := sharedInt
		if val%10 == 0 {
			unusedValue++
		}
	})
}

func runSimpleWriter() {
	for {
		mutex.Lock()
		sharedInt = sharedInt + 1
		mutex.Unlock()
		//fmt.Println(sharedInt)
	}
}

func main() {
	go runSimpleReader()
	go runSimpleWriter()

	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println(sharedInt, unusedValue)
	time.Sleep(2 * time.Second)
}
