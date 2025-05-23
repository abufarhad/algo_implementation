package main

import (
	"fmt"
	"sync"
)

var cnt = 0

func inc(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	ch <- true
	cnt++
	<-ch
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan bool, 2)
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		inc(&wg, ch)
	}

	wg.Wait() /**/

	fmt.Println("Final value of cnt = ", cnt)
}
