name: Madi-Engine-Core CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build-and-test:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout Repository
        uses: actions/checkout@v4

      - name: Set up Rust Toolchain
        uses: actions-rs/toolchain@v1
        with:
          toolchain: stable
          override: true

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Install Protoc
        run: |
          sudo apt-get update
          sudo apt-get install -y protobuf-compiler

      - name: Build and Test Rust Core
        run: |
          cd core/rust-core
          cargo build --verbose
          cargo test --verbose

      - name: Build and Test Go Gateway
        run: |
          cd apps/gateway
          go get google.golang.org/grpc
          go get google.golang.org/grpc/credentials/insecure
          go mod tidy
          cd ../..
          go build -v ./apps/gateway/...
          go test -v ./apps/gateway/...
