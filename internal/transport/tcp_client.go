package transport

import (
	"fmt"
	"net"

	"github.com/akansha204/mini-rpc/internal/protocol"
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

	err = protocol.WriteFrame(conn, []byte(message))

	if err != nil {
		return err
	}

	// buffer := make([]byte, 1024)
	payload, err := protocol.ReadFrame(conn)

	if err != nil {
		return err
	}
	fmt.Println("Echoed response:", string(payload))
	return nil
}
