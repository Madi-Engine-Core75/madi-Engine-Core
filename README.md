# Madi-Engine-Core

Madi-Engine-Core is a high-performance, secure middleware kernel designed as an orchestration layer for financial routing and decentralized communication modules.

## Architecture & Tech Stack

* **Core (Rust - 74.7%):** Manages secure state evaluation, cryptographic operations (AES-256-GCM), and internal vault storage.
* **Gateway (Go - 25.3%):** Handles high-concurrency network traffic, routing protocols, and account authentication (`/api/v1/auth/login`).
* **CI/CD Pipeline:** Automated build and test workflows configured via GitHub Actions.

## Repository Structure

* `src/`: Core Rust logic and storage encryption modules.
* `gateway/`: Go high-concurrency routing and authentication gateway.
* `.github/workflows/`: Automated CI/CD integration pipelines.

