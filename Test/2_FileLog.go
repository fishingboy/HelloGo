package main

import (
	"log"
	"os"
)

func main() {
	FileLog("test 1111")
	FileLog("test 2222")
	FileLog("test 3333")
}

func FileLog(message string) bool {
	// 打開檔案，若檔案不存在則創建
	file, err := os.OpenFile("D:\\app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("無法打開或創建 log 檔案: %v", err)
	}
	defer file.Close()

	// 設置 log 輸出到檔案
	log.SetOutput(file)

	// 寫入 log
	log.Println(message)

	return true
}
