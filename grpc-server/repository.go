package main

import (
	pb "grpc.example/proto"
	"sync"
)

type repository interface {
	Create(*pb.Todo) (*pb.Todo, error)
	GetAll() []*pb.Todo
}

// Repository - Dummy repository, this simulates the use of a data store
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	mu    sync.RWMutex
	todos []*pb.Todo
}

func (repo *Repository) Create(todo *pb.Todo) (*pb.Todo, error) {
	repo.mu.Lock()
	updated := append(repo.todos, todo)
	repo.todos = updated
	repo.mu.Unlock()
	return todo, nil
}

func (repo *Repository) GetAll() []*pb.Todo {
	return repo.todos
}
