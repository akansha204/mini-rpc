package protocol

import "encoding/json"

type Request struct {
	ID      uint64          `json:"id"`
	Method  string          `json:"method"`
	Payload json.RawMessage `json:"payload"`
}

type Response struct {
	ID      uint64          `json:"id"`
	Payload json.RawMessage `json:"payload"`
	Error   string          `json:"error"`
}
