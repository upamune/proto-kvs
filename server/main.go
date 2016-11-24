package main

import (
	"log"
	"net"
	"sync"

	context "golang.org/x/net/context"

	pb "github.com/upamune/proto-kvs/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type kvsServer struct {
	store map[string]string
	mu    sync.RWMutex
}

func (s *kvsServer) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if val, ok := s.store[r.Key]; ok {
		return &pb.GetResponse{
			Value: val,
		}, nil
	}
	return &pb.GetResponse{}, grpc.Errorf(codes.NotFound, "key not found")
}

func (s *kvsServer) Set(ctx context.Context, r *pb.SetRequest) (*pb.SetResponse, error) {
	s.mu.Lock()
	s.store[r.Key] = r.Value
	s.mu.Unlock()

	return &pb.SetResponse{}, nil
}

func (s *kvsServer) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	s.mu.Lock()
	delete(s.store, r.Key)
	s.mu.Unlock()

	return &pb.DeleteResponse{}, nil
}

func (s *kvsServer) List(ctx context.Context, r *pb.ListRequest) (*pb.ListResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return &pb.ListResponse{
		Store: s.store,
	}, nil
}

func NewKvsServer() *kvsServer {
	return &kvsServer{
		store: make(map[string]string),
	}
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterKvsServer(s, NewKvsServer())

	s.Serve(l)
}
