package main

import (
	"context"
	"fmt"
	"time"
)

// 模擬一個需要獨立執行的操作，使用 context.Background()
func independentOperation(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Independent operation completed.")
	case <-ctx.Done():
		fmt.Println("Independent operation canceled.") // 不會被觸發，因為 ctx 不可取消
	}
}

// 使用傳入的 context 來控制一個操作
func controlledOperation(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Controlled operation completed.")
	case <-ctx.Done():
		fmt.Println("Controlled operation canceled.") // 會在 context 被取消時觸發
	}
}

func main() {
	// 創建一個可以取消的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 啟動獨立的操作，與 ctx 無關
	go independentOperation(context.Background())

	// 啟動受控的操作，受 ctx 控制
	go controlledOperation(ctx)

	// 模擬主程式運行 2 秒後取消受控的操作
	time.Sleep(2 * time.Second)
	fmt.Println("Main function: canceling controlled operation...")
	cancel()

	// 等待一會兒，以便所有操作完成
	time.Sleep(6 * time.Second)
	fmt.Println("Main function: done.")
}
