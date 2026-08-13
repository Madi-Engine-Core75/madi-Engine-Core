package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Madi-Engine-Core75/madi-gateway/internal/router"
	pb "github.com/Madi-Engine-Core75/madi-gateway/proto/gen/core"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to rust-core: %v", err)
	}
	defer conn.Close()

	client := pb.NewMadiEngineCoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	healthRes, err := client.HealthCheck(ctx, &pb.HealthRequest{})
	if err != nil {
		log.Printf("HealthCheck failed: %v", err)
	} else {
		log.Printf("Rust Core Status: %s (Timestamp: %d)", healthRes.Status, healthRes.Timestamp)
	}

	r := router.NewRouter()
	log.Println("Madi Gateway is up and running, routing events to MadiEngineCore.")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run gateway server: %v", err)
	}
}
