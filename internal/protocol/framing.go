package protocol

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
)

const headerSize = 4
const maxFrameSize = 4 * 1024 * 1024

func WriteFrame(c net.Conn, payload []byte) error {
	length := len(payload)
	if length > maxFrameSize {
		return errors.New("Payload too large")
	}
	header := make([]byte, headerSize)
	binary.BigEndian.PutUint32(header, uint32(length))
	if _, err := c.Write(header); err != nil {
		return err
	}
	if _, err := c.Write(payload); err != nil {
		return err
	}
	return nil
}

func ReadFrame(c net.Conn) ([]byte, error) {
	header := make([]byte, headerSize)
	if _, err := io.ReadFull(c, header); err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(header)
	if length > maxFrameSize {
		return nil, errors.New("Frame too large")
	}
	payload := make([]byte, length)
	if _, err := io.ReadFull(c, payload); err != nil {
		return nil, err
	}
	return payload, nil
}
