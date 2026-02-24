package main

import (
	"log"

	"github.com/akansha204/mini-rpc/internal/transport"
)

func main() {
	client := transport.NewTCPClient(":8080")
	err := client.Send("meowww from client")
	if err != nil {
		log.Fatal(err)
	}
}
