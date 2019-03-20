package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "exercises/JSON_GRPC_microservice/api/proto/v1"
	mapdb "exercises/JSON_GRPC_microservice/server/pkg"

	"google.golang.org/grpc"
)

const (
	port = "127.0.0.1:50051"
)

type server struct{}

func (s *server) GetPort(ctx context.Context, in *pb.GetPortRequest) (*pb.GetPortReply, error) {
	log.Printf("Received request GetPort: %v", in.Portkey)
	return &pb.GetPortReply{Message: in.Portkey + " " + fmt.Sprintf("%s", mapdb.GetPort(in.Portkey))}, nil
}
func (s *server) AddUpdatePort(ctx context.Context, in *pb.AddUpdatePortRequest) (*pb.AddUpdatePortReply, error) {
	log.Printf("Received request AddUpdatePort: %v, %s", in.Key, in.Data)
	return &pb.AddUpdatePortReply{Message: mapdb.AddUpdatePort(in.Key, in.Data)}, nil
}

func (s *server) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckReprly, error) {
	log.Printf("Received request HealthCheck: %v", "")
	if mapdb.IsAlive() {
		return &pb.HealthCheckReprly{Status: pb.HealthCheckReprly_SERVING}, nil
	}
	return &pb.HealthCheckReprly{Status: pb.HealthCheckReprly_NOT_SERVING}, nil
}

func main() {
	if !mapdb.ConnectToDB() {
		log.Panic("Can not connect to DB")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPortDomainServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
