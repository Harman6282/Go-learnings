package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/Harman6282/grpc/pb/proto"

	"google.golang.org/grpc"
)

type taskServer struct {
	pb.UnimplementedTaskManagerServer
	mu      sync.Mutex
	tasks   []*pb.Task
	counter int
}

func (s *taskServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++
	id := fmt.Sprintf("task-%d", s.counter)

	task := &pb.Task{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
	}
	s.tasks = append(s.tasks, task)

	log.Printf("Created task: %s", id)

	return &pb.CreateTaskResponse{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
	}, nil
}

func (s *taskServer) ListTasks(req *pb.ListTaksRequest, stream pb.TaskManager_ListTasksServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, task := range s.tasks {
		if err := stream.Send(task); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Listen on port 50051 — this is the conventional default port for gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server and register our TaskManager implementation
	grpcServer := grpc.NewServer()
	pb.RegisterTaskManagerServer(grpcServer, &taskServer{})

	log.Println("gRPC server is running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}