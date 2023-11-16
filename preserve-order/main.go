package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // wait for 1 go routine
	go task1(&wg)
	wg.Wait()

	wg.Add(3) // wait for 1 go routine
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
	wg.Wait()

	wg.Add(1) // wait for 1 go routine
	go task5(&wg)
	wg.Wait()
	fmt.Println("done waiting, main exits")
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done task1")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done task2")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done task3")
}

func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done task4")
}

func task5(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done task5")
}
