package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/DariusChandra/grpc-gateway-experiment/gen/go/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

type server struct {
	pb.UnimplementedYourServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Received: %v", in.GetValue())
	return &pb.StringMessage{Value: in.GetValue()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
