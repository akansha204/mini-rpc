package main

import (
	"log"

	"github.com/akansha204/mini-rpc/internal/codec"
	"github.com/akansha204/mini-rpc/internal/transport"
	"github.com/akansha204/mini-rpc/rpc"

	user "github.com/akansha204/mini-rpc/demo/user"
)

func main() {
	jsonCodec := &codec.JSONCodec{}

	server := rpc.NewServer(
		jsonCodec,
		rpc.NewRegistry(),
	)

	user.RegisterUserService(
		server,
		&UserServiceImpl{},
	)

	tcpServer := transport.NewTCPServer(
		":8080",
		server.Handle,
	)

	log.Println("Server listening on :8080")

	if err := tcpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
