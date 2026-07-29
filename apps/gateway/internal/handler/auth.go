package handler

import (
"gateway/proto/gen/core"
)

type Handler struct {
core.UnimplementedAuthServiceServer
}

func NewHandler() *Handler {
return &Handler{}
}
