package messages

import "encoding/json"

const (
	// Message types
	InitType        = "init"
	InitOkType      = "init_ok"
	EchoType        = "echo"
	EchoOkType      = "echo_ok"
	GenerateType    = "generate"
	GenerateOkType  = "generate_ok"
	BroadcastType   = "broadcast"
	BroadcastOkType = "broadcast_ok"
	ReadType        = "read"
	ReadOkType      = "read_ok"
	TopologyType    = "topology"
	TopologyOkType  = "topology_ok"

	// Keys in the messages - these might become less necessary if using specific structs
	EchoKey      = "echo"
	NodeIDKey    = "node_id"
	TypeKey      = "type"
	MsgIDKey     = "msg_id"
	InReplyToKey = "in_reply_to"
	IDKey        = "id"
	MessageKey   = "message"
	MessagesKey  = "messages"
	NodeIDsKey   = "node_ids" // Added for consistency with RequestBody
	TopologyKey  = "topology" // Added for consistency with RequestBody
)

type Message struct {
	Src  string          `json:"src"`
	Dest string          `json:"dest"`
	Body json.RawMessage `json:"body"`
}

func NewMessage(src string, dest string, body any) (Message, error) {
	marshaledBody, err := json.Marshal(body)
	if err != nil {
		return Message{}, err
	}
	return Message{src, dest, marshaledBody}, nil
}

type RequestBody struct {
	NodeID    string              `json:"node_id"`
	NodeIds   []string            `json:"node_ids"`
	Echo      string              `json:"echo"`
	Type      string              `json:"type"`
	MsgID     int                 `json:"msg_id"`
	InReplyTo int                 `json:"in_reply_to"`
	Message   int                 `json:"message"`
	Topology  map[string][]string `json:"topology"`
}

type EchoResponseBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
	Echo      string `json:"echo"`
}

type InitResponseBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
}

type GenerateResponseBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
	ID        string `json:"id"`
}

type BroadcastResponseBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
}

type ReadResponseBody struct {
	Type      string `json:"type"`
	Messages  []int  `json:"messages"`
	InReplyTo int    `json:"in_reply_to"`
}

type TopologyResponseBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
}
