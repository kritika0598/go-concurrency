package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var count int64
	wg.Add(1)
	// reader
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Count in go routine", atomic.LoadInt64(&count))
	}()

	wg.Add(50)
	// writers
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()
	fmt.Println("count in main", count)
}
