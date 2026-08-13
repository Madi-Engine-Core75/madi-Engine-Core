printf "module github.com/Madi-Engine-Core75/madi-gateway\n\ngo 1.26.4\n" > go.mod


go 1.26.4

replace github.com/Madi-Engine-Core75/madi-gateway/proto/gen/core => ../../proto/gen/core
replace github.com/Madi-Engine-Core75/madi-gateway/internal/client => ../../internal/client
replace github.com/Madi-Engine-Core75/madi-gateway/internal/vault => ../../internal/vault
replace github.com/Madi-Engine-Core75/madi-gateway/internal/handler => ../../internal/handler
replace github.com/Madi-Engine-Core75/madi-gateway/internal/router => ../../internal/router
