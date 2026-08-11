package client

import (
	"context"
	"fmt"
	"time"

	pb "./proto/vault"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type VaultGRPCClient struct {
	client pb.VaultServiceClient
	conn   *grpc.ClientConn
}

func NewVaultGRPCClient(targetAddr string) (*VaultGRPCClient, error) {
	// ملاحظة: في بيئة الإنتاج الصارمة، يتم استبدال insecure بـ TLS certificates (mTLS)
	conn, err := grpc.Dial(targetAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rust core: %v", err)
	}

	client := pb.NewVaultServiceClient(conn)
	return &VaultGRPCClient{client: client, conn: conn}, nil
}

func (v *VaultGRPCClient) Close() {
	v.conn.Close()
}

func (v *VaultGRPCClient) Encrypt(plaintext []byte) ([]byte, []byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &pb.EncryptRequest{Plaintext: plaintext}
	res, err := v.client.EncryptData(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	if res.Error != "" {
		return nil, nil, fmt.Errorf("core error: %s", res.Error)
	}

	return res.Ciphertext, res.Nonce, nil
}
