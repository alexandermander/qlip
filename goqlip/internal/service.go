package service


import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

type store struct {
    mu     sync.RWMutex
    buffer string
}

// globalStore holds our data in memory (for demonstration)
var globalStore = store{}

// Your "secret" pass (for demonstration only).
// In production, this shouldn't be hardcoded (e.g. read from env variables).
const secretPass = "QQ1122ww"

func main() {
    // For demonstration, log that the server has started
    fmt.Println("Server is starting on port :8080...")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World! Send data to /up with ?qlip=...&pass=...")
    })

    // Handler that stores data
    http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {

        pass := r.URL.Query().Get("pass")
        if pass != secretPass {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        qlipData := r.URL.Query().Get("qlip")
        if qlipData == "" {
            http.Error(w, "Missing qlip parameter", http.StatusBadRequest)
            return
        }

        // 3. Store it safely (thread-safe)
        globalStore.mu.Lock()
        globalStore.buffer = qlipData
        globalStore.mu.Unlock()

        fmt.Fprintln(w, "Data stored successfully!")
    })

    // Handler that retrieves stored data
    http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
        // Optionally require pass as well
        pass := r.URL.Query().Get("pass")
        if pass != secretPass {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Retrieve safely
        globalStore.mu.RLock()
        stored := globalStore.buffer
        globalStore.mu.RUnlock()

        // Return the data
        fmt.Fprintf(w, "Currently stored: %s\n", stored)
    })

    // NOTE: Typically we do http.ListenAndServe(":8080", nil) for local dev.
    // If you need HTTPS on port 443 in production, see TLS section below.
    // For local testing, use 8080 or 3000, etc., or behind a reverse proxy on Render.

    log.Fatal(http.ListenAndServe(":8080", nil))
}

