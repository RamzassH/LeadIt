package main

import (
	"context"
	authv1 "github.com/RamzassH/LeadIt/gateway/contracts/contracts/gen/auth"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	grpcServerAddress := os.Getenv("GRPC_SERVER_ADDRESS")
	if grpcServerAddress == "" {
		grpcServerAddress = "localhost:50051"
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := authv1.RegisterAuthHandlerFromEndpoint(ctx, mux, grpcServerAddress, opts)

	if err != nil {
		log.Fatal("Не удалось зарегистрировать gateway: %v", err)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Printf("Запуск HTTP gateway на порту %s. gRPC сервер: %s", httpPort, grpcServerAddress)
	// Запускаем HTTP сервер
	if err := http.ListenAndServe(":"+httpPort, mux); err != nil {
		log.Fatalf("Ошибка при запуске HTTP сервера: %v", err)
	}
}
