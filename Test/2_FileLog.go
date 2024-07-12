package main

import (
	"log"
	"os"
)

func main() {
	// 打開檔案，若檔案不存在則創建
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("無法打開或創建 log 檔案: %v", err)
	}
	defer file.Close()

	// 設置 log 輸出到檔案
	log.SetOutput(file)

	// 寫入 log
	log.Println("這是一條資訊 log")
	log.Println("這是一條警告 log")
	log.Println("這是一條錯誤 log")
}
