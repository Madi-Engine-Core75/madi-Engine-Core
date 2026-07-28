.PHONY: all proto rust-build gateway-build test clean

all: proto rust-build gateway-build

# توليد ملفات gRPC للـ Go و Rust
proto:
	@echo "==> Generating Protocol Buffers..."
	protoc --go_out=. --go-grpc_out=. proto/vault.proto

# بناء النواة الأمنية (Rust)
rust-build:
	@echo "==> Building Rust Core..."
	cd core/rust-core && cargo build --release

# بناء بوابة الـ Go (Gateway)
gateway-build:
	@echo "==> Building Go Gateway..."
	cd apps/gateway && go build -o bin/gateway ./cmd/main.go

# تنفيذ الاختبارات
test:
	@echo "==> Running tests across modules..."
	cd core/rust-core && cargo test
	cd apps/gateway && go test ./...

# تنظيف الملفات المؤقتة
clean:
	@echo "==> Cleaning build artifacts..."
	rm -rf apps/gateway/bin/
	cd core/rust-core && cargo clean
