package main

import (
	"github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/server"
	"log"
)

func main() {
	server := server.CreateServer(":5001")
	log.Println("Starting HTTP server on port 5001")
	server.ListenAndServe()
}
