package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/upamune/proto-kvs/proto"
	"google.golang.org/grpc"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	c := pb.NewKvsClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			os.Exit(0)
		}

		args := strings.Split(scanner.Text(), " ")

		switch args[0] {
		case "get":
			if len(args) < 1 {
				continue
			}
			var key = args[1]
			resp, err := c.Get(context.Background(), &pb.GetRequest{Key: key})
			if err != nil {
				fmt.Println("key not found")
				continue
			}
			fmt.Println(resp.Value)
		case "set":
			if len(args) < 3 {
				fmt.Println("needs three arguments")
				continue
			}
			var key = args[1]
			var value = args[2]
			_, err := c.Set(context.Background(), &pb.SetRequest{Key: key, Value: value})
			if err != nil {
				continue
			}
			fmt.Println("OK")
		case "del":
			if len(args) < 1 {
				continue
			}
			var key = args[1]
			_, err := c.Delete(context.Background(), &pb.DeleteRequest{Key: key})
			if err != nil {
				continue
			}
			fmt.Println("OK")
		case "list":
			if len(args) < 1 {
				continue
			}
			resp, err := c.List(context.Background(), &pb.ListRequest{})
			if err != nil {
				continue
			}

			for key, val := range resp.Store {
				fmt.Printf("Key: [%s] Value: [%s]\n", key, val)
			}
		case "exit":
			os.Exit(0)
		default:
			if scanner.Text() == "" {
				continue
			}
			fmt.Println("Unknown command")
		}
	}
}
