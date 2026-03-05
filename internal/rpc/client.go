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
func (c *Client) Call(method string, params []interface{}, result interface{}) error {
	c.id++

	req := protocol.Request{
		ID:     c.id,
		Method: method,
		Params: params,
	}

	data, err := c.codec.Encode(req)
	if err != nil {
		return err
	}

	err = c.conn.Send(data)
	if err != nil {
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
	data, _ = c.codec.Encode(resp.Result)
	return c.codec.Decode(data, result)

}
