package user

import (
	"github.com/akansha204/mini-rpc/rpc"
)

type UserRequest struct {
	Name string
	Age  int32
}

type UserResponse struct {
	Name string
	Age  int32
}

type UserService interface {
	GetUser(req UserRequest) (UserResponse, error)
	CreateUser(req UserRequest) (UserResponse, error)
}

type UserServiceClient struct {
}

func NewUserServiceClient() *UserServiceClient {
	return &UserServiceClient{}
}

func (c *UserServiceClient) GetUser(req UserRequest) (UserResponse, error) {
	panic("not implemented")
}

func (c *UserServiceClient) CreateUser(req UserRequest) (UserResponse, error) {
	panic("not implemented")
}

func RegisterUserService(server *rpc.Server, service UserService) {

	server.Register(
		"UserService/GetUser",
		func(payload []byte) ([]byte, error) {

			var req UserRequest

			if err := server.Decode(payload, &req); err != nil {
				return nil, err
			}

			resp, err := service.GetUser(req)
			if err != nil {
				return nil, err
			}

			return server.Encode(resp)
		},
	)

	server.Register(
		"UserService/CreateUser",
		func(payload []byte) ([]byte, error) {

			var req UserRequest

			if err := server.Decode(payload, &req); err != nil {
				return nil, err
			}

			resp, err := service.CreateUser(req)
			if err != nil {
				return nil, err
			}

			return server.Encode(resp)
		},
	)
}
