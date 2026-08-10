package handler

import ("madi-gateway/proto/gen/core")

type Handler struct {core.UnimplementedAuthServiceServer}

func NewHandler() *Handler {return &Handler{}}
