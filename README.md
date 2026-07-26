# Madi-Engine-Core

High-performance, secure middleware kernel designed as an orchestration layer for financial routing and decentralized communication modules.

## Architecture & Tech Stack

* **Core (`core/rust-core/`):** Built with Rust, manages secure state evaluation, cryptographic operations (AES-256-GCM), and internal vault storage.
* **Gateway (`apps/gateway/`):** Built with Go, handles high-concurrency network traffic, routing protocols, and account authentication (`/api/v1/auth/login`).
* **CI/CD Pipeline (`.github/workflows/ci.yml`):** Automated build, test, and verification workflows.

## Repository Structure

```text
Madi-Engine-Core/
├── apps/
│   └── gateway/          # Go gateway and API handlers
├── core/
│   └── rust-core/        # Rust core logic and cryptographic modules
├── proto/                # Shared gRPC contracts
├── .github/workflows/    # CI/CD automation pipelines
└── Makefile              # Unified build automation
