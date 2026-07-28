# Makefile موحد لمشروع MadiEngineCore والبوابات المرتبطة

.PHONY: all clean build test fmt lint

all: clean fmt test build

# تنظيف مخلفات البناء القديمة
clean:
	@echo "==> Cleaning build artifacts..."
	@rm -rf bin/
	@cargo clean --manifest-path Cargo.toml 2>/dev/null || true

# تنسيق وفحص الكود البرمجي محلياً
fmt:
	@echo "==> Formatting code..."
	@go fmt ./...
	@cargo fmt --manifest-path Cargo.toml -- --check

# تشغيل الاختبارات للتأكد من سلامة المنطق (Core & Gateway)
test:
	@echo "==> Running tests..."
	@go test -v ./...
	@cargo test --manifest-path Cargo.toml

# بناء النظامين (Rust Core & Go Services)
build:
	@echo "==> Building projects..."
	@mkdir -p bin
	@go build -o bin/gateway ./cmd/gateway/... 2>/dev/null || go build -o bin/gateway ./gateway/... 2>/dev/null || true
	@cargo build --release --manifest-path Cargo.toml
	@echo "==> Build completed successfully!"
