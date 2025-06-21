package main

import (
	"log"
	"net/http"

	"github.com/bag-huyag/api-gateway/internal/client"
	"github.com/bag-huyag/api-gateway/internal/handler"
	"github.com/gorilla/mux"
	// "github.com/bag-huyag/api-gateway/pkg/server"
)

func main() {
	// if err := server.StartGRPC(); err != nil {
	// 	log.Fatalf("Не удалось запустить gRPC сервер", err)
	// }

	userClient := client.NewUserServiceClient("localhost:50052")

	h := handler.Handler{UserClient: userClient}

	r := mux.NewRouter()

	r.HandleFunc("/users", h.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")

	// http.HandleFunc("/users", h.CreateUser)

	log.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
