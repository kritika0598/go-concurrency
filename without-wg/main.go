package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("number of cores", runtime.NumCPU())
	now := time.Now()
	for i := 0; i < 10; i++ {
		go work(i)
	}
	time.Sleep(time.Second)
	fmt.Println("time elapsed:", time.Since(now))
	fmt.Println("main is done")
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "is done")
}
