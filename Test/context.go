package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 創建一個帶有取消功能的 context
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("[main] start workers ...")

	// 啟動一個 goroutine，並將 context 傳遞給它
	go doWork(ctx, "worker1")
	go doWork(ctx, "worker2")
	go doWork(ctx, "worker3")

	// 讓主程序等待 2 秒
	time.Sleep(2 * time.Second)

	// 呼叫取消函數來取消 context
	fmt.Println("[main] Requesting to cancel the work...")
	cancel()

	// 再等待 1 秒，確保 goroutine 已經退出
	time.Sleep(1 * time.Second)
	fmt.Println("[main] Main function exits.")
}

func doWork(ctx context.Context, workName string) {
	// 模擬一個需要不斷執行的工作
	for {
		select {
		case <-ctx.Done():
			// 當 context 被取消時，Done 通道會被關閉
			fmt.Printf("[%s] Work canceled!\n", workName)
			return
		default:
			// 沒有取消信號時，持續工作
			fmt.Printf("[%s] Working...\n", workName)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
