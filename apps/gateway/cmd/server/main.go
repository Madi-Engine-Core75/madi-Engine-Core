package main

import (
	context "context"
	log "log"
	net "net"
	time "time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
pb "github.com/Madi-Engine-Core75/madi-gateway/internal/pb"
	func main() {
	// 1. الاتصال بخادم Rust (gRPC Server)
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to rust-core: %v", err)
	}
	defer conn.Close()

	client := pb.NewMadiEngineCoreClient(conn)

	// 2. اختبار فحص الحالة (HealthCheck) مع النواة
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	healthRes, err := client.HealthCheck(ctx, &pb.HealthRequest{})
	if err != nil {
		log.Printf("HealthCheck failed: %v", err)
	} else {
		log.Printf("Rust Core Status: %s (Timestamp: %d)", healthRes.Status, healthRes.Timestamp)
	}

	// 3. هنا يمكنك لاحقاً ربط مسارات الـ API (مثل /api/v1/auth/login) لتمرير البيانات عبر اللتلخيص والتشفير
	log.Println("Madi Gateway is up and running, routing events to MadiEngineCore.")
}
