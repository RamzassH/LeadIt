package main

import (
	"context"
	authv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/auth"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	grpcServerAddress := os.Getenv("GRPC_SERVER_ADDRESS")
	if grpcServerAddress == "" {
		grpcServerAddress = "localhost:57442"
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			md := metadata.New(map[string]string{})
			if cookies := req.Cookies(); len(cookies) > 0 {
				var cookieStrings []string
				for _, c := range cookies {
					cookieStrings = append(cookieStrings, c.String())
				}
				md.Set("cookie", strings.Join(cookieStrings, "; "))
			}
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("Ошибка %s", err)
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
		}),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := authv1.RegisterAuthHandlerFromEndpoint(ctx, mux, grpcServerAddress, opts)
	if err != nil {
		log.Fatalf("Не удалось зарегистрировать gateway: %v", err)
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
	}).Handler(mux)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Printf("Запуск HTTP gateway на порту %s. gRPC сервер: %s", httpPort, grpcServerAddress)
	if err := http.ListenAndServe(":"+httpPort, corsHandler); err != nil {
		log.Fatalf("Ошибка при запуске HTTP сервера: %v", err)
	}
}
