package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup, workerName string) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
		fmt.Printf("[%s] run..\n", workerName)
	}
}

func main() {
	// 競態條件

	var wg sync.WaitGroup
	wg.Add(5)

	go increment(&wg, "worker1")
	go increment(&wg, "worker2")
	go increment(&wg, "worker3")
	go increment(&wg, "worker4")
	go increment(&wg, "worker5")

	wg.Wait()
	fmt.Println("Final counter2 value:", counter)
}
