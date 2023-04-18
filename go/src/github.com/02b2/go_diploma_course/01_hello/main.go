package main

import (
	"fmt"
	"time"
)

func main() {
	delay := time.Second * 2

	fmt.Println("Go: Hi, who are you?")
	time.Sleep(delay)

	fmt.Println("Dero: I'm Dero.")
	time.Sleep(delay)

	fmt.Println("Go: HELLO DERO!!! nice to meet you!!")
	time.Sleep(delay)

	fmt.Println("Dero: Nice to meet you too Go.")
}
