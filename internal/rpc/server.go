package rpc

import (
	"github.com/akansha204/mini-rpc/internal/codec"
	"github.com/akansha204/mini-rpc/internal/protocol"
)

type Server struct {
	codec    codec.Codec
	registry *Registry
}

func NewServer(c codec.Codec, r *Registry) *Server {
	return &Server{
		codec:    c,
		registry: r,
	}
}

func (s *Server) Handle(data []byte) ([]byte, error) {
	var req protocol.Request

	if err := s.codec.Decode(data, &req); err != nil {
		return nil, err
	}

	handler, ok := s.registry.Get(req.Method)
	if !ok {
		resp := protocol.Response{
			ID:    req.ID,
			Error: "method not found",
		}
		return s.codec.Encode(resp)
	}

	result, err := handler(req.Params)
	if err != nil {
		resp := protocol.Response{
			ID:    req.ID,
			Error: err.Error(),
		}
		return s.codec.Encode(resp)
	}

	resp := protocol.Response{
		ID:     req.ID,
		Result: result,
	}
	return s.codec.Encode(resp)
}
