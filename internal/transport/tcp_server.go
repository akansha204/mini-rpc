package transport

import (
	"fmt"
	"io"
	"net"

	"github.com/akansha204/mini-rpc/internal/protocol"
)

type TCPServer struct {
	addr string
}

func NewTCPServer(addr string) *TCPServer {
	return &TCPServer{
		addr: addr,
	}
}

func (s *TCPServer) Start() error {
	li, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	defer li.Close()

	fmt.Println("Server listening on", s.addr)

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		go s.handleConnection(conn)
	}
}
func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New connection from", conn.RemoteAddr())

	// buffer := make([]byte, 1024)

	for {
		payload, err := protocol.ReadFrame(conn)
		if err == io.EOF {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		fmt.Println("Received:", string(payload))

		err = protocol.WriteFrame(conn, payload) //writing back to the tcp
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}
