package main

import (
	"fmt"
	"log"

	"github.com/akansha204/mini-rpc/internal/codec"
	"github.com/akansha204/mini-rpc/internal/rpc"
	"github.com/akansha204/mini-rpc/internal/transport"
)

func main() {
	client := transport.NewTCPClient(":8080")
	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	jsonCodec := &codec.JSONCodec{}
	rpcClient := rpc.NewClient(jsonCodec, client)
	var result int

	err := rpcClient.Call("Add", []interface{}{5, 3}, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result", result)

	err = rpcClient.Call("Subtract", []interface{}{5, 3}, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result", result)

	err = rpcClient.Call("Mul", []interface{}{5, 3}, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result", result)

	err = rpcClient.Call("Div", []interface{}{5, 5}, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result", result)

}
