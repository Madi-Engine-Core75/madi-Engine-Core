.PHONY: all build-core build-gateway clean test

# الأمر الافتراضي لبناء الكل
all: clean build-core build-gateway

# بناء نواة Rust
build-core:
	@echo "==> Building Rust Core..."
	cd core/rust-core && cargo build --release

# بناء بوابة Golang
build-gateway:
	@echo "==> Building Go Gateway..."
	cd apps/gateway && go build -o bin/gateway cmd/server/main.go

# تنظيف ملفات البناء المؤقتة
clean:
	@echo "==> Cleaning up build artifacts..."
	rm -rf apps/gateway/bin
	cd core/rust-core && cargo clean

# تشغيل الاختبارات
test:
	@echo "==> Running tests..."
	cd core/rust-core && cargo test
	cd apps/gateway && go test ./...
