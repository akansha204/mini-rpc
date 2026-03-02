package protocol

type Request struct {
	ID     uint64        `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type Response struct {
	ID     uint64      `json:"id"`
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}
