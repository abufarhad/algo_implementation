package main

import (
	"sync/atomic"
	"time"
)

// If the shared data is a single variable, we can use counters provided in the sync/atomic pack
// Instead of accessing the shared variable directly, we can use the atomic.LoadInt64()/atomic.AddInt64() pair to access it.

var sharedInt int64 = 0
var unusedValue int = 0

func runSimpleReader() {
	val := atomic.LoadInt64(&sharedInt)
	if val%10 == 0 {
		unusedValue++
	}
}

func runSimpleWriter() {
	for {
		atomic.AddInt64(&sharedInt, 1)
	}
}

func main() {
	go runSimpleReader()
	go runSimpleWriter()
	time.Sleep(10 * time.Second)
}
