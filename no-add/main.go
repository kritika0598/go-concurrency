package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("number of cores", runtime.NumCPU())
	wg := sync.WaitGroup{}
	now := time.Now()
	go func() {
		defer wg.Done()
		fmt.Println("go routine: Done")
	}()
	wg.Wait()
	fmt.Println("time elapsed:", time.Since(now))
	fmt.Println("main is done")
}
