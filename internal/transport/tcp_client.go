package transport

import (
	"errors"
	"fmt"
	"net"

	"github.com/akansha204/mini-rpc/internal/protocol"
)

type TCPClient struct {
	addr string
	conn net.Conn
}

func NewTCPClient(addr string) *TCPClient {
	return &TCPClient{
		addr: addr,
	}
}

func (c *TCPClient) Connect() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *TCPClient) Send(message string) error {
	if c.conn == nil {
		return errors.New("not connected")
	}

	err := protocol.WriteFrame(c.conn, []byte(message))

	if err != nil {
		return err
	}

	// buffer := make([]byte, 1024)
	payload, err := protocol.ReadFrame(c.conn)

	if err != nil {
		return err
	}
	fmt.Println("Echoed response:", string(payload))
	return nil
}

func (c *TCPClient) Close() error {
	if c.conn != nil {
		c.conn.Close()
	}
	return nil
}
