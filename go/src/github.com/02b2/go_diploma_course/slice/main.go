package main

import (
	"fmt"
	"time"
)

func main() {
	myArray := []string{"Apple", "Banana", "Cherry", "Date", "Fig", "Grape"}

	if len(myArray) >= 6 {
		for _, element := range myArray {
			fmt.Println(element)
			time.Sleep(500 * time.Millisecond)
		}
	} else {
		fmt.Println("The array must have at least 6 elements.")
	}
}
