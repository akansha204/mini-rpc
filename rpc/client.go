package rpc

import (
	"errors"

	"github.com/akansha204/mini-rpc/internal/codec"
	"github.com/akansha204/mini-rpc/internal/protocol"
	"github.com/akansha204/mini-rpc/internal/transport"
)

type Client struct {
	codec codec.Codec
	conn  *transport.TCPClient
	id    uint64
}

func NewClient(c codec.Codec, conn *transport.TCPClient) *Client {
	return &Client{
		codec: c,
		conn:  conn,
	}
}

func NewDefaultClient(addr string) (*Client, error) {
	conn := transport.NewTCPClient(addr)
	if err := conn.Connect(); err != nil {
		return nil, err
	}
	return &Client{
		codec: &codec.JSONCodec{},
		conn:  conn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
func (c *Client) Call(method string, req interface{}, result interface{}) error {
	c.id++

	payload, err := c.codec.Encode(req)
	if err != nil {
		return err
	}

	request := protocol.Request{
		ID:      c.id,
		Method:  method,
		Payload: payload,
	}

	data, err := c.codec.Encode(request)
	if err != nil {
		return err
	}

	if err := c.conn.Send(data); err != nil{
		return err
	}
	respBytes, err := c.conn.Receive()
	if err != nil {
		return err
	}

	var resp protocol.Response
	err = c.codec.Decode(respBytes, &resp)
	if err != nil {
		return err
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return c.codec.Decode(resp.Payload, result)

}
