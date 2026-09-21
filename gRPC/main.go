package main

import (
	"log"
	"net"

	pb "grpc/basic/proto"

	"google.golang.org/grpc"
)



func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("tcp connection failed")
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, nil)

	log.Println("grpc server started at port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
