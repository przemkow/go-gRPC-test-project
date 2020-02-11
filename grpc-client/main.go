package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc.example/proto"
	"io"
	"log"
	"os"
	"strings"
)

const (
	address = "localhost:50051"
)

func addTodo(client pb.TodoServiceClient) {
	fmt.Println("-----------------")
	fmt.Println("| Add new todo: |")
	fmt.Println("-----------------")

	reader := bufio.NewReader(os.Stdin)
	todoName, _ := reader.ReadString('\n')
	todoName = strings.TrimSuffix(todoName, "\n")

	todo := pb.Todo{Done: false, Name: todoName}

	r, err := client.AddTodo(context.Background(), &todo)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	fmt.Printf("New Task Created! \nName: '%s'", r.Todo.Name)
	fmt.Printf("\n\n")
}

func listTodos(client pb.TodoServiceClient) {
	stream, err := client.GetTodos(context.Background(), &pb.NoArgs{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	fmt.Println("----------")
	fmt.Println("| TODOs: |")
	fmt.Println("----------")

	for {
		todo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		fmt.Printf("Name: %s \nDone: %t \n\n", todo.Name, todo.Done)
	}
	fmt.Printf("\n\n")
}
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTodoServiceClient(conn)



	mainLoop:for {
		fmt.Println("Select action:")
		fmt.Println("(1) Create todo")
		fmt.Println("(2) List todos")
		fmt.Println("(exit) Exit program")

		reader := bufio.NewReader(os.Stdin)
		action, _ := reader.ReadString('\n')
		action = strings.TrimSuffix(action, "\n")
		switch action {
		case "1":
			addTodo(client)
		case "2":
			listTodos(client)
		case "exit":
			break mainLoop
		default:
			fmt.Printf("Incorrect action: %s.\n", action)
		}
	}



}