package main

import (
	"context"
	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

var (
	grpcServerEndpoint = os.Getenv("GRPC_SERVER_ADDRESS")
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := sl.RegisterLoyaltyServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatal("failed to register proxy handler: %v", err)
	}

	httpServer := &http.Server{
		Addr:    ":50051",
		Handler: mux,
	}
	log.Print("Starting HTTP server at :50051")
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start HTTP server: %v", err)
	}
}
