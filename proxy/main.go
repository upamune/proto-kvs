package main

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/upamune/proto-kvs/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	mux := runtime.NewServeMux()

	err := pb.RegisterKvsHandlerFromEndpoint(context.Background(), mux, "localhost:50051", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", mux)
}
