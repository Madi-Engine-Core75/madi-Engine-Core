package handler

import ("./proto/gen/core")

type Handler struct {core.UnimplementedAuthServiceServer}

func NewHandler() *Handler {return &Handler{}}
