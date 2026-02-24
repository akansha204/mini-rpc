package transport

import (
	"fmt"
	"net"
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

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		fmt.Println("Received:", string(buffer[:n]))

		_, err = conn.Write(buffer[:n]) //writing back to the tcp
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}
