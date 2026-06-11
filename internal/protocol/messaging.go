package protocol

type Request struct {
	ID      uint64 `json:"id"`
	Method  string `json:"method"`
	Payload []byte `json:"payload"`
}

type Response struct {
	ID      uint64 `json:"id"`
	Payload []byte `json:"payload"`
	Error   string `json:"error"`
}
