package main

import (
	"fmt"
	"time"
)

func main() {
	oneSecond := make(chan string)
	fiveSeconds := make(chan string)

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Second)
			oneSecond <- "One second"
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Second * 5)
			fiveSeconds <- "Five seconds"
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case <-oneSecond:
			fmt.Println(<-oneSecond)
		case <-fiveSeconds:
			fmt.Println(<-fiveSeconds)
		}
	}

}
