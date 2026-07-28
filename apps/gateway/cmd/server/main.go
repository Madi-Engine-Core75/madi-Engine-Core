package main

import (
	"fmt"
	"net/http"
)

func handleGatewayRoute(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Madi-Token")
	
	if token == "" {
		http.Error(w, "Unauthorized: Missing security token", http.StatusUnauthorized)
		return
	}

	// محاكاة توجيه الطلب بنجاح نحو MadiEngineCore
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Gateway routing passed: Request securely forwarded to MadiEngineCore.")
}

func main() {
	http.HandleFunc("/api/v1/route", handleGatewayRoute)
	
	fmt.Println("Madi-Gateway is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Gateway server failed: %v\n", err)
	}
}
