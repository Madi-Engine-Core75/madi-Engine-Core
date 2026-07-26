package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/MadiEngine-Core75/Madi-Engine-Core/proto/gen/core"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token,omitempty"`
	Error  string `json:"error,omitempty"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w.Encode(LoginResponse{Status: "error", Error: "Invalid request payload"}))
		return
	}

	// الاتصال بالنواة الأمنية عبر gRPC (يُفترض تشغيل النواة على المنفذ 50051 محلياً)
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w.Encode(LoginResponse{Status: "error", Error: "Failed to connect to security core"}))
		return
	}
	defer conn.Close()

	client := pb.NewVaultServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// إرسال طلب المصادقة للنواة
	res, err := client.AuthenticateUser(ctx, &pb.AuthRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w.Encode(LoginResponse{Status: "error", Error: "Authentication core error"}))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w.Encode(LoginResponse{
		Status: res.Status,
		Token:  res.Token,
		Error:  res.Error,
	}))
}
