package main

import (
"log"
"gateway/internal/router"
"gateway/internal/handler"
)

func main() {
log.Println("Starting gateway server...")
r := router.NewRouter()
_ = r
_ = handler.NewHandler()
log.Println("Gateway server initialized successfully.")
}
