package main

import (
	"fmt"
	"time"
)

func main() {
	go work() // a fork point is created
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done waiting, main exits")
	// Joint point should be here
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing work")
}
