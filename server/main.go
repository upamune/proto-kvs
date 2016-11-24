package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	context "golang.org/x/net/context"

	pb "github.com/upamune/proto-kvs/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type kvsServer struct {
	store map[string]string
	mu    sync.RWMutex
	c     map[chan pb.Entry]struct{}
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
	defer s.mu.Unlock()
	s.store[r.Key] = r.Value

	for c := range s.c {
		c <- pb.Entry{Key: r.Key, Value: r.Value}
	}

	return &pb.SetResponse{}, nil
}

func (s *kvsServer) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	s.mu.Lock()
	delete(s.store, r.Key)
	s.mu.Unlock()

	for c := range s.c {
		c <- pb.Entry{Key: r.Key}
	}

	return &pb.DeleteResponse{}, nil
}

func (s *kvsServer) List(ctx context.Context, r *pb.ListRequest) (*pb.ListResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return &pb.ListResponse{
		Store: s.store,
	}, nil
}

func (s *kvsServer) Watch(r *pb.WatchRequest, ws pb.Kvs_WatchServer) error {
	ch := make(chan pb.Entry)

	s.mu.Lock()
	s.c[ch] = struct{}{}
	s.mu.Unlock()
	fmt.Println("Added New Watcher", ch)

	defer func() {
		s.mu.Lock()
		delete(s.c, ch)
		s.mu.Unlock()
		close(ch)
		fmt.Println("Deleted Watcher", ch)
	}()

	for c := range ch {
		if !strings.HasPrefix(c.Key, r.Prefix) {
			continue
		}
		err := ws.Send(&c)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewKvsServer() *kvsServer {
	return &kvsServer{
		store: make(map[string]string),
		c:     make(map[chan pb.Entry]struct{}),
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
