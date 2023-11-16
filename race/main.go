package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var count int32
	wg.Add(5)
	go func() {
		defer wg.Done()
		count += 10
	}()
	go func() {
		defer wg.Done()
		count = 15
	}()
	go func() {
		defer wg.Done()
		count++
	}()
	go func() {
		defer wg.Done()
		count = 0
	}()
	go func() {
		defer wg.Done()
		count = 100
	}()
	wg.Wait()
	fmt.Println("count: ", count)
}
