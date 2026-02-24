package transport

import (
	"fmt"
	"net"
)

type TCPClient struct {
	addr string
}

func NewTCPClient(addr string) *TCPClient {
	return &TCPClient{
		addr: addr,
	}
}

func (c *TCPClient) Send(message string) error {
	conn, err := net.Dial("tcp", c.addr)

	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))

	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		return err
	}
	fmt.Println("Echoed response:", string(buffer[:n]))
	return nil
}
