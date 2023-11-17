package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var count int32
	wg.Add(5)
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&count, 10)
	}()
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&count, -15)
	}()
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&count, 1)
	}()
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&count, 0)
	}()
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&count, 100)
	}()
	wg.Wait()
	fmt.Println("count: ", count)
}
