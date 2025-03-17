package main

import (
    "fmt"
    "log"
	"github.com/yoshioka0101/spot_map/internal/server"
)

func main() {
    fmt.Println("Hello, World!!!!!Test Air")

	r := server.NewRouter()

	// サーバーの起動
	log.Println("Starting server on :8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
