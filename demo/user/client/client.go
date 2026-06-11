package main

import (
	"fmt"
	"log"

	"github.com/akansha204/mini-rpc/internal/codec"
	"github.com/akansha204/mini-rpc/internal/transport"
	"github.com/akansha204/mini-rpc/rpc"

	user "github.com/akansha204/mini-rpc/demo/user"
)

func main() {
	tcpClient := transport.NewTCPClient(":8080")

	if err := tcpClient.Connect(); err != nil {
		log.Fatal(err)
	}
	defer tcpClient.Close()

	client := rpc.NewClient(
		&codec.JSONCodec{},
		tcpClient,
	)

	req := user.UserRequest{
		Name: "Akansha",
		Age:  21,
	}

	var resp user.UserResponse

	err := client.Call(
		"UserService/GetUser",
		req,
		&resp,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}
