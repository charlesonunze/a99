package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"

	"github.com/charlesonunze/a99/internal/handler"
	"github.com/charlesonunze/a99/internal/model"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/charlesonunze/a99/pb"
	"github.com/spf13/viper"
)

func main() {
	GRPC_PORT := os.Getenv("GRPC_PORT")
	G8WAY_PORT := os.Getenv("G8WAY_PORT")

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	db.AutoMigrate(&model.Car{}, &model.Feature{})

	s := grpc.NewServer()
	server := handler.New(db)
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
		Handler: cors(gwmux),
	}

	log.Println("Serving gRPC-Gateway on", G8WAY_PORT)
	log.Fatalln(gwServer.ListenAndServe())
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
