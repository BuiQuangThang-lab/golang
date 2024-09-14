package main

import (
	"google.golang.org/grpc"
	"net"
	"qlnv/internal/intergation"
	"qlnv/internal/presentation"
	"qlnv/pkg/database"
	pb "qlnv/userpb"
)

func main() {
	db := database.GetMySQLInstance()
	presentation.StartHTTPServer(db)
	s := grpc.NewServer()
	userService := intergation.NewUserService()
	pb.RegisterUserServiceServer(s, userService)
	lis, _ := net.Listen("tcp", ":8888")
	err := s.Serve(lis)
	if err != nil {
		return
	}
}
