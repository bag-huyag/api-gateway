package main

import (
	"log"

	"github.com/bag-huyag/api-gateway/pkg/server"
)

func main() {
	if err := server.StartGRPC(); err != nil {
		log.Fatalf("Не удалось запустить gRPC сервер", err)
	}
}
