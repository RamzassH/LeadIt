package main

import (
	"context"
	authv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/auth"
	organizationv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/organization"
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
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			md := metadata.New(map[string]string{})

			authHeader := req.Header.Get("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
				md.Set("authorization", tokenStr)
			}

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

	services := []struct {
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
		address      string
	}{
		{authv1.RegisterAuthHandlerFromEndpoint, os.Getenv("AUTH_GRPC_SERVER_ADDRESS")},
		{organizationv1.RegisterOrganizationHandlerFromEndpoint, os.Getenv("ORGANIZATION_GRPC_SERVER_ADDRESS")},
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}

	for _, service := range services {
		if err := service.registerFunc(ctx, mux, service.address, opts); err != nil {
			log.Fatalf("Failed to register gRPC service: %v", err)
		}
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
	}).Handler(mux)

	httpPort := os.Getenv("GATEWAY_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Printf("Starting HTTP gateway port %s", httpPort)
	if err := http.ListenAndServe(":"+httpPort, corsHandler); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
