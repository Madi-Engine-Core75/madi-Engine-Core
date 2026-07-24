package main

import (
"bufio"
"fmt"
"net"
"time"
)

func main() {
fmt.Println("🚀 Nadi Engine Gateway is running and connecting to Core...")
for {
conn, err := net.Dial("tcp", "127.0.0.1:50051")
if err != nil {
time.Sleep(2 * time.Second)
continue
}

reader := bufio.NewReader(conn)
for {
message, err := reader.ReadString('^')
if err != nil {
fmt.Println("⚠️ Disconnected from Engine Core, reconnecting...")
break
}
fmt.Printf("📥 Sync Data Received: %s\n", message)
}
conn.Close()
}
}
