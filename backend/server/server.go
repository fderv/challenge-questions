package main

import (
	"context"
	"net"
	"log"
	"strconv"
	"encoding/json"
	"google.golang.org/grpc"
	"github.com/go-redis/redis/v8"
	pb "go-grpc-server/user"
)

var rClient *redis.Client
var ctx = context.Background()

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.User) (*pb.Reply, error) {
	//log.Printf("Received: %v", in.GetEmail())
	out, err := json.Marshal(in)
	if err != nil {
        panic (err)
	}
	rClient.RPush(ctx, "userdata", string(out))
	num := rClient.LLen(ctx, "userdata")
	return &pb.Reply{ReplMsg: "Added " + in.GetEmail() + " Total data: " + strconv.FormatInt(num.Val(), 10)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	rClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	
	
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
