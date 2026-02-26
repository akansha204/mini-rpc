package main

import (
	"log"

	"github.com/akansha204/mini-rpc/internal/transport"
)

func main() {
	client := transport.NewTCPClient(":8080")
	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	err := client.Send("meowww from client!\nYOKOSO PIRAA PIRAA PUKKI PUKKKI WAKU WAKU WOSHAIIII")
	err = client.Send("YOWAIII MOOO!!!")
	err = client.Send("ARAARRAA GOMEENNN!!!")
	if err != nil {
		log.Fatal(err)
	}
}
