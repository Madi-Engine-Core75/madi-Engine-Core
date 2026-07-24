# Madi-Engine-Core

## Overview
`madi-engine-core` is the central middleware kernel designed to bridge and synchronize standalone applications within the ecosystem. It acts as the core orchestration layer connecting financial transaction routing and decentralized social communication modules.

---

## Technical Stack & Architecture
* **Rust (52.1%):** High-performance core logic, secure vault protocols, and memory-safe processing modules.
* **Python (30.3%):** Integration scripts, utility services, and automation pipelines.
* **Go (17.6%):** High-concurrency gateway handling and network transport layers.

---

## Repository Structure
```text
madi-Engine-Core/
├── cargo/          # Rust workspace configurations and crates
├── gateway/        # High-concurrency Go routing services
├── src/            # Core source logic and cross-language bridges
├── .gitignore      # Explicit exclusion rules for build artifacts
├── Cargo.lock      # Rust dependency lock file
├── Cargo.toml      # Rust package manifest
└── main.py         # Primary entry point and orchestration script
