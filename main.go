package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/charlesonunze/a99/internal/handler"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/charlesonunze/a99/pb"
)

func main() {
	GRPC_PORT := os.Getenv("GRPC_PORT")
	G8WAY_PORT := os.Getenv("G8WAY_PORT")

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	server := handler.New()
	pb.RegisterCarServiceServer(s, server)

	// Serve gRPC server
	log.Println("Serving gRPC on", GRPC_PORT)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterCarServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    G8WAY_PORT,
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on", G8WAY_PORT)
	log.Fatalln(gwServer.ListenAndServe())
}
