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
	var msgBody RequestBody
	if err := json.Unmarshal(msg.Body, &msgBody); err != nil {
		return Message{}, err
	}
	switch msgBody.Type {
	case initType:
		return getReplyToInit(msg, node)
	case echoType:
		return getReplyToEcho(msg)
	case generateType:
		return getReplyToGenerate(msg)
	case broadcastType:
		return getReplyToBroadcast(msg, node)
	case readType:
		return getReplyToRead(msg, node)
	case topologyType:
		return getReplyToTopology(msg, node)
	default:
		return Message{}, fmt.Errorf("unknown message type: %s", msgBody.Type)
	}
}
func getReplyToInit(msg Message, node *Node) (Message, error) {
	var requestBody RequestBody
	json.Unmarshal(msg.Body, &requestBody)
	node.NodeID = requestBody.NodeId
	node.NodeIds = requestBody.NodeIds

	return NewMessage(
		msg.Dest,
		msg.Src,
		InitResponseBody{
			initOkType,
			requestBody.MsgId,
		})
}

func getReplyToEcho(msg Message) (Message, error) {
	var requestBody RequestBody
	json.Unmarshal(msg.Body, &requestBody)

	return NewMessage(
		msg.Dest,
		msg.Src,
		EchoResponseBody{
			echoOkType,
			requestBody.MsgId,
			requestBody.Echo,
		})
}

func getReplyToGenerate(msg Message) (Message, error) {
	var requestBody RequestBody
	json.Unmarshal(msg.Body, &requestBody)

	id := strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + msg.Dest
	return NewMessage(
		msg.Dest,
		msg.Src,
		GenerateResponseBody{
			generateOkType,
			requestBody.MsgId,
			id,
		})
}

func getReplyToBroadcast(msg Message, node *Node) (Message, error) {
	var requestBody RequestBody
	json.Unmarshal(msg.Body, &requestBody)

	node.Messages = append(node.Messages, requestBody.Message)

	return NewMessage(
		msg.Dest,
		msg.Src,
		BroadcastResponseBody{
			broadcastOkType,
			requestBody.MsgId,
		},
	)
}

func getReplyToRead(msg Message, node *Node) (Message, error) {
	var requestBody RequestBody
	json.Unmarshal(msg.Body, &requestBody)

	return NewMessage(
		msg.Dest,
		msg.Src,
		ReadResponseBody{
			readOkType,
			node.Messages,
			requestBody.MsgId,
		},
	)
}

func getReplyToTopology(msg Message, node *Node) (Message, error) {
	var requestBody RequestBody
	json.Unmarshal(msg.Body, &requestBody)
	node.Topology = requestBody.Topology

	return NewMessage(
		msg.Dest,
		msg.Src,
		TopologyResponseBody{
			topologyOkType,
			requestBody.MsgId,
		},
	)
}
