package main

import "encoding/json"

const (
	// Message types
	initType        = "init"
	initOkType      = "init_ok"
	echoType        = "echo"
	echoOkType      = "echo_ok"
	generateType    = "generate"
	generateOkType  = "generate_ok"
	broadcastType   = "broadcast"
	broadcastOkType = "broadcast_ok"
	readType        = "read"
	readOkType      = "read_ok"
	topologyType    = "topology"
	topologyOkType  = "topology_ok"

	// Keys in the messages
	echoKey      = "echo"
	nodeIdKey    = "node_id"
	typeKey      = "type"
	msgIdKey     = "msg_id"
	inReplyToKey = "in_reply_to"
	idKey        = "id"
	messageKey   = "message"
	messagesKey  = "messages"
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
	NodeId    string              `json:"node_id"`
	NodeIds   []string            `json:"node_ids"`
	Echo      string              `json:"echo"`
	Type      string              `json:"type"`
	MsgId     int                 `json:"msg_id"`
	InReplyTo int                 `json:"in_reply_to"`
	Message   any                 `json:"message"`
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
	Id        string `json:"id"`
}

type BroadcastResponseBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
}

type ReadResponseBody struct {
	Type      string `json:"type"`
	Messages  []any  `json:"messages"`
	InReplyTo int    `json:"in_reply_to"`
}

type TopologyResponseBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
}
