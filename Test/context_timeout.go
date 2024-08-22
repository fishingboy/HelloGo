package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	// 設置一個 2 秒超時的 context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 確保在 main 函數結束時取消 context
	fmt.Println(time.Now())

	// 模擬一個需要執行的操作
	resultChan := make(chan string)
	fmt.Println(time.Now())

	go func() {
		// 模擬耗時操作，例如網絡請求或計算
		time.Sleep(3 * time.Second)
		resultChan <- "Operation completed"
	}()

	fmt.Println(time.Now())
	select {
	case result := <-resultChan:
		fmt.Println(result)
	case <-ctx.Done():
		// 超時處理
		fmt.Println("Operation timed out:", ctx.Err())
	}
	fmt.Println(time.Now())
}

func run() {
	// 模擬耗時操作，例如網絡請求或計算
	time.Sleep(3 * time.Second)
	resultChan <- "Operation completed"
}
