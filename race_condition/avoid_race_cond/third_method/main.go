package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var sharedMapForMutex map[string]int = map[string]int{}

// RWMutex performed more read operations
// But please keep in mind that this example is highly biased as there are 15 reader goroutines and a single writer goroutine

var mapMutex = sync.RWMutex{}

//var mapMutex = sync.Mutex{}

func runMapMutexReader(ctx context.Context, readCh chan int) {
	readCount := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("reader exiting and readCount: ", readCount)
			readCh <- readCount
			return

		default:
			mapMutex.Lock()
			_ = sharedMapForMutex["key"]
			mapMutex.Unlock()
			readCount++
		}
	}
}

func runMapMutexWriter(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writer exiting")
			return

		default:
			mapMutex.Lock()
			sharedMapForMutex["key"]++
			mapMutex.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	testContext, cancel := context.WithCancel(context.Background())
	readCh := make(chan int)
	sharedMapForMutex["key"] = 0

	numberOfReaders := 15

	for i := 0; i < numberOfReaders; i++ {
		go runMapMutexReader(testContext, readCh)
	}
	go runMapMutexWriter(testContext)
	time.Sleep(2 * time.Second)
	cancel()

	totalReadCount := 0
	for i := 0; i < numberOfReaders; i++ {
		totalReadCount += <-readCh
	}

	time.Sleep(1 * time.Second)

	counter := sharedMapForMutex["key"]
	fmt.Printf("[MUTEX] Write Counter value: %v\n", counter)
	fmt.Printf("[MUTEX] Read Counter value: %v\n", totalReadCount)
}
