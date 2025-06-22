package main

import (
	"bytes"
	"strings"
)

const (
	messageTemplate  = `{"src":"%+v","dest":"%+v","body":%+v}`
	initBodyTemplate = `{
		"type":"init",
		"msg_id":%+v,
		"node_id":"%+v",
		"node_ids":["n1", "n2", "n3"]
	}`
	initOkBodyTemplate = `{
		"type":"init_ok",
		"in_reply_to":%+v
	}`
	errorBodyTemplate = `{
		"type":"error",
		"in_reply_to":5,
		"code":11,
		"text":"Node n5 is waiting for quorum and cannot service requests yet"
	}`
	echoBodyTemplate = `{
		"type":"echo",
		"msg_id":%+v,
		"echo":"%+v"
	}`
	echoOkBodyTemplate = `{
		"type":"echo_ok",
		"in_reply_to":%+v,
		"echo":"%+v"
	}`
	unknownTypeBody = `{
		"type":"unknown_type"
	}`
	generateBodyTemplate = `{
		"type":"generate",
		"msg_id":%+v
	}`
	generateOkBodyTemplate = `{
		"type": "generate_ok",
		"in_reply_to":%+v,
		"id":"%+v"
	}`
)

func newUnknownTypeMessage(src string, dest string) (Message, error) {
	body := RequestBody{Type: unknownTypeBody}
	return NewMessage(src, dest, body)
}

func newInitMessage(src string, dest string, msgId int, nodeId string) (Message, error) {
	body := RequestBody{Type: initType, MsgId: msgId, NodeId: nodeId}
	return NewMessage(src, dest, body)
}

func newInitOkMessage(src string, dest string, inReplyTo int) (Message, error) {
	body := InitResponseBody{initOkType, inReplyTo}
	return NewMessage(src, dest, body)
}

func newEchoMessage(src string, dest string, msgId int, echo string) (Message, error) {
	body := RequestBody{Type: echoType, MsgId: msgId, Echo: echo}
	return NewMessage(src, dest, body)
}

func newEchoOkMessage(src string, dest string, inReplyTo int, echo string) (Message, error) {
	body := EchoResponseBody{echoOkType, inReplyTo, echo}
	return NewMessage(src, dest, body)
}

func newGenerateMessage(src string, dest string, msgId int) (Message, error) {
	body := RequestBody{Type: generateType, MsgId: msgId}
	return NewMessage(src, dest, body)
}

func newGenerateOkMessage(src string, dest string, msgId int, id string) (Message, error) {
	body := GenerateResponseBody{generateOkType, msgId, id}
	return NewMessage(src, dest, body)
}

func newTestNode(input string) Node {
	in := strings.NewReader(input)
	out := new(bytes.Buffer)

	return Node{
		In:  in,
		Out: out,
	}
}
