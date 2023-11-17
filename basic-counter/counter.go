package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup
var count int = 0

func incCount() {
	count++
	wg.Done()
}

func incCountWithMutex() {
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

func doCount() {
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go incCount()
	}
}

func doCountWithMutex() {
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go incCountWithMutex()
	}
}

func main() {
	count = 0
	doCount()
	wg.Wait()
	fmt.Println("Count without mutex: ", count)

	count = 0
	doCountWithMutex()
	wg.Wait()
	fmt.Println("Count with mutex: ", count)
}
