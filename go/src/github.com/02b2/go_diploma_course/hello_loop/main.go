package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 6; i++ {
		fmt.Println("Hello Dero")
		time.Sleep(1 * time.Second)
	}
}
