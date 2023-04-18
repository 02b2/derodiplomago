package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(msg string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(delay)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printMessage("Hello from Goroutine 1", 1*time.Second, &wg)
	go printMessage("Hello from Goroutine 2", 2*time.Second, &wg)

	wg.Wait()
}
