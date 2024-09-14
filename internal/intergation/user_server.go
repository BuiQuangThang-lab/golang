package intergation

import (
	"context"
	log "github.com/sirupsen/logrus"
	"qlnv/internal/dto"
	"qlnv/internal/service"
	"qlnv/userpb"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewUserService() *UserServiceServer {
	return &UserServiceServer{}
}

func (u *UserServiceServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	log.Println("Creating user: %v\n", req)

	// Thực hiện logic tạo user ở đây
	userDTO := dto.UserDTO{
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		FullName: req.FullName,
		Role:     int(req.Role),
		Gender:   int(req.Gender),
		Birthday: req.Birthday,
		Password: req.Password,
	}
	log.Println(userDTO)
	err := u.userService.CreateUser(userDTO)
	if err != nil {
		return &userpb.CreateUserResponse{
			Status: "500",
			Des:    "User created fail",
		}, nil
	}
	return &userpb.CreateUserResponse{
		Status: "200",
		Des:    "User created successfully",
	}, nil
}
