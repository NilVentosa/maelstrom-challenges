package main

import "encoding/json"

type Message struct {
	Src  string          `json:"src"`
	Dest string          `json:"dest"`
	Body json.RawMessage `json:"body"`
}

func newMessage(src string, dest string, body any) (Message, error) {
	marshaledBody, err := json.Marshal(body)
	if err != nil {
		return Message{}, err
	}
	return Message{src, dest, marshaledBody}, nil
}

type RequestBody struct {
	NodeId    string `json:"node_id"`
	Echo      string `json:"echo"`
	Type      string `json:"type"`
	MsgId     int    `json:"msg_id"`
	InReplyTo int    `json:"in_reply_to"`
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
