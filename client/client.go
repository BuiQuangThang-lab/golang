package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "qlnv/userpb"
)

func main() {
	// Kết nối đến gRPC server
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Tạo yêu cầu CreateUserRequest
	req := &pb.CreateUserRequest{
		UserName: "JohnDoe",
		Email:    "john.doe@example.com",
		Phone:    "1234567890",
		FullName: "John Doe",
		Role:     1,
		Gender:   1,
		Birthday: "1990-01-01",
		Password: "securepassword",
	}

	// Gọi API CreateUser
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("req", req)
	resp, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}

	// In kết quả
	log.Printf("CreateUser Response: %s", resp.GetDes())
}
