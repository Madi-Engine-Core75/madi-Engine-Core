package handler

import ("github.com/Madi-Engine-Core75/madi-gateway/proto/gen/core")

type Handler struct {core.UnimplementedAuthServiceServer}

func NewHandler() *Handler {return &Handler{}}
