package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// 在本地地址和端口 8080 上監聽 TCP 連接
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on localhost:8080...")

	for {
		// 接受來自客戶端的連接
		// 這行會一直停住，直接下一個 user 連進來！
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 處理客戶端連接
		// 開一個 goroutine 來處理
		// 都是丟進 goroutine ，所以允許多個連線
		go handleConnection(conn)
	}
}

// 處理客戶端連接的函數
func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	// 使用 bufio 來讀取客戶端的輸入
	reader := bufio.NewReader(conn)
	for {
		// 讀取來自客戶端的一行輸入
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}

		// 去掉字串末尾的換行符號
		message = strings.TrimSpace(message)
		fmt.Printf("Received message: %s\n", message)

		// 回送消息給客戶端
		_, err = conn.Write([]byte("Echo: " + message + "\r\n"))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}

		// 如果客戶端發送 "quit"，則關閉連接
		if message == "quit" {
			fmt.Println("Client requested to close the connection.")
			return
		}
	}
}
