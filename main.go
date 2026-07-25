package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	// تشغيل خادم الـ HTTP الخاص بالمصادقة والواجهات على المنفذ 8080 في الخلفية
	go func() {
		http.HandleFunc("/api/v1/auth/login", handleAccountAuth)
		fmt.Println("HTTP Gateway Server is listening on :8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("HTTP Server Error: %v\n", err)
		}
	}()

	fmt.Println("Madi Engine Gateway is running and connecting to Core...")
	for {
		conn, err := net.Dial("tcp", "127.0.0.1:5001")
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}

		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("⚠️ Disconnected from Engine Core, reconnecting...")
				break
			}
			fmt.Printf("⚠️ Sync Data Received: %s\n", message)
		}
		conn.Close()
	}
}

