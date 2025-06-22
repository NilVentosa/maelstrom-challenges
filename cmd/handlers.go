package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func handleMessage(msg Message, node *Node) ([]byte, error) {
	response, responseError := getReplyToMessage(msg, node)

	if responseError != nil {
		return nil, responseError
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("problem marshaling the response: %w", err)
	}

	return jsonResponse, nil
}

func getReplyToMessage(msg Message, node *Node) (Message, error) {
	var body RequestBody
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return Message{}, err
	}
	switch body.Type {
	case initType:
		return getReplyToInit(msg, body, node)
	case echoType:
		return getReplyToEcho(msg, body)
	case generateType:
		return getReplyToGenerate(msg, body)
	case broadcastType:
		return getReplyToBroadcast(msg, body, node)
	case readType:
		return getReplyToRead(msg, body, node)
	case topologyType:
		return getReplyToTopology(msg, body, node)
	default:
		return Message{}, fmt.Errorf("unknown message type: %s", body.Type)
	}
}
func getReplyToInit(msg Message, body RequestBody, node *Node) (Message, error) {
	node.NodeID = body.NodeId
	node.NodeIds = body.NodeIds

	return NewMessage(
		msg.Dest,
		msg.Src,
		InitResponseBody{
			initOkType,
			body.MsgId,
		})
}

func getReplyToEcho(msg Message, body RequestBody) (Message, error) {
	return NewMessage(
		msg.Dest,
		msg.Src,
		EchoResponseBody{
			echoOkType,
			body.MsgId,
			body.Echo,
		})
}

func getReplyToGenerate(msg Message, body RequestBody) (Message, error) {
	id := strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + msg.Dest
	return NewMessage(
		msg.Dest,
		msg.Src,
		GenerateResponseBody{
			generateOkType,
			body.MsgId,
			id,
		})
}

func getReplyToBroadcast(msg Message, body RequestBody, node *Node) (Message, error) {
	node.Messages = append(node.Messages, body.Message)

	return NewMessage(
		msg.Dest,
		msg.Src,
		BroadcastResponseBody{
			broadcastOkType,
			body.MsgId,
		},
	)
}

func getReplyToRead(msg Message, body RequestBody, node *Node) (Message, error) {
	return NewMessage(
		msg.Dest,
		msg.Src,
		ReadResponseBody{
			readOkType,
			node.Messages,
			body.MsgId,
		},
	)
}

func getReplyToTopology(msg Message, body RequestBody, node *Node) (Message, error) {
	node.Topology = body.Topology

	return NewMessage(
		msg.Dest,
		msg.Src,
		TopologyResponseBody{
			topologyOkType,
			body.MsgId,
		},
	)
}
