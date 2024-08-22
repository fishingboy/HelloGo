package main

import (
	"fmt"
	"sync"
)

var counter2 int
var mu sync.Mutex

func increment2(wg *sync.WaitGroup, workerName string) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mu.Lock() // 鎖定資源
		counter2++
		fmt.Printf("[%s] run..\n", workerName)
		mu.Unlock() // 釋放資源
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	go increment2(&wg, "worker1")
	go increment2(&wg, "worker2")
	go increment2(&wg, "worker3")
	go increment2(&wg, "worker4")
	go increment2(&wg, "worker5")

	wg.Wait()
	fmt.Println("Final counter2 value:", counter2)
}
