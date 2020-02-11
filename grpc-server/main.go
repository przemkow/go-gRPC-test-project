package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc.example/proto"
)

const (
	port = ":50051"
)


// Service should implement all the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface.
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

func (s *service) AddTodo(ctx context.Context, req *pb.Todo) (*pb.Response, error) {
	todo, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	log.Println("new todo saved")
	return &pb.Response{Created: true, Todo: todo}, nil
}

func (s *service) GetTodos(_ *pb.NoArgs, stream pb.TodoService_GetTodosServer) error {
	for _, todo := range s.repo.GetAll() {
		log.Println(todo)
		if err := stream.Send(todo); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	repo := &Repository{}


	//Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Running on port:", port)

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &service{repo})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
