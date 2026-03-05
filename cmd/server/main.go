package main

import (
	"fmt"
	"log"

	"github.com/akansha204/mini-rpc/internal/codec"
	"github.com/akansha204/mini-rpc/internal/rpc"
	"github.com/akansha204/mini-rpc/internal/transport"
)

func main() {
	registry := rpc.NewRegistry()
	calc := &rpc.Calculator{}

	registry.Register("Add", calc.Add)
	registry.Register("Subtract", calc.Subtract)
	registry.Register("Mul", calc.Mul)
	registry.Register("Div", calc.Div)

	jsonCodec := &codec.JSONCodec{}
	rpcServer := rpc.NewServer(jsonCodec, registry)
	fmt.Println("mini-rpc Server Starting...")
	tcpserver := transport.NewTCPServer(":8080", rpcServer.Handle)
	if err := tcpserver.Start(); err != nil {
		log.Fatal(err)
	}
}
