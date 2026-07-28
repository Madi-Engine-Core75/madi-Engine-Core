package router

import (
"log"
)

type Router struct{}

func NewRouter() *Router {
log.Println("Router initialized")
return &Router{}
}
