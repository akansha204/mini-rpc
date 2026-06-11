package main

import "github.com/akansha204/mini-rpc/demo/user"

type UserServiceImpl struct{}

func (s *UserServiceImpl) GetUser(req user.UserRequest) (user.UserResponse, error) {
	return user.UserResponse{
		Name: req.Name,
		Age:  req.Age,
	}, nil
}

func (s *UserServiceImpl) CreateUser(
	req user.UserRequest,
) (user.UserResponse, error) {

	return user.UserResponse{
		Name: req.Name,
		Age:  req.Age,
	}, nil
}
