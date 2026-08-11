cat << 'EOF' > go.mod
module github.com/Madi-Engine-Core75/madi-gateway

go 1.26.4

replace github.com/Madi-Engine-Core75/madi-gateway => ./
replace github.com/Madi-Engine-Core75/MadiEngineCore => ../../rust-core
replace github.com/Madi-Engine-Core75/madi-gateway/proto/gen/core => ./internal/proto/gen/core
replace github.com/Madi-Engine-Core75/madi-gateway/proto/vault => ./internal/proto/vault
EOF

go mod tidy
git add go.mod
git commit -m "fix: point proto replaces to internal directory path"
git push origin fix/workflow-and-configs-update
