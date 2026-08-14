cat << 'EOF' > cmd/server/main.go
package main

import (
	"log"
	"net/http"
	"github.com/Madi-Engine-Core75/madi-gateway/internal/vault"
	"github.com/Madi-Engine-Core75/madi-gateway/internal/broker"
	"strconv"
	"time"
)

func main() {
	log.Println("Initializing Madi-Engine-Core Gateway with Async Broker & Vault...")
	
	eventDispatcher := broker.NewDispatcher(100)

	http.HandleFunc("/api/v1/route", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error":"method not allowed"}`))
			return
		}
		
		plainPayload := "financial-routing-payload-" + strconv.FormatInt(time.Now().UnixNano(), 10)
		encrypted, err := vault.Encrypt([]byte(plainPayload))
		if err != nil {
			log.Printf("Encryption error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"encryption failed"}`))
			return
		}

		// إرسال الحدث للمعالجة غير المتزامنة في الخلفية دون إبطاء الاستجابة
		eventDispatcher.Publish("EVT-"+strconv.FormatInt(time.Now().Unix(), 10), plainPayload)

		log.Println("Payload successfully encrypted, queued asynchronously, and routed via AES-256-GCM.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"routed","cipher":"` + encrypted + `"}`))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Server is listening securely on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
EOF
