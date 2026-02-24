package main

import (
	"fmt"
	"log"

	"github.com/akansha204/mini-rpc/internal/transport"
)

func main() {
	fmt.Println("mini-rpc Server Starting...")
	server := transport.NewTCPServer(":8080")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	server.Start()
}
