package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type FileLine struct {
	Filename string
	Line     string
}

func printFileContent(filename string, ch chan FileLine, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- FileLine{Filename: filename, Line: scanner.Text()}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch1 := make(chan FileLine)
	ch2 := make(chan FileLine)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go printFileContent("file1.txt", ch1, wg)
	go printFileContent("file2.txt", ch2, wg)

	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	delayBetweenFiles := 4 * time.Second

	type chanFile struct {
		channel  chan FileLine
		fileLine FileLine
		ok       bool
	}

	cf1 := chanFile{channel: ch1}
	cf2 := chanFile{channel: ch2}

	for {
		cf1.fileLine, cf1.ok = <-cf1.channel
		if cf1.ok {
			fmt.Printf("%s: %s\n", cf1.fileLine.Filename, cf1.fileLine.Line)
			time.Sleep(delayBetweenFiles)
		}

		cf2.fileLine, cf2.ok = <-cf2.channel
		if cf2.ok {
			fmt.Printf("%s: %s\n", cf2.fileLine.Filename, cf2.fileLine.Line)
			time.Sleep(delayBetweenFiles)
		}

		if !cf1.ok && !cf2.ok {
			break
		}
	}
}
