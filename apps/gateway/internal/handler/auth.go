package handler

import (
"github.com/MadiEngine-Core75/madi-Engine-Core/apps/gateway/proto/gen/core"
)

type Handler struct {
core.UnimplementedAuthServiceServer
}

func NewHandler() *Handler {
return &Handler{}
}
