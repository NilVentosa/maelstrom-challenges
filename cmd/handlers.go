package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"time"
)

func handleMessage(msg Message, node *Node) error {
	var body RequestBody
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	switch body.Type {
	case initType:
		return handleInit(msg, body, node)
	case echoType:
		return handleEcho(msg, body, node)
	case generateType:
		return handleGenerate(msg, body, node)
	case broadcastType:
		return handleBroadcast(msg, body, node)
	case broadcastOkType:
		return nil
	case readType:
		return handleRead(msg, body, node)
	case topologyType:
		return handleTopology(msg, body, node)
	default:
		return fmt.Errorf("unknown message type: %s", body.Type)
	}
}
func handleInit(msg Message, body RequestBody, node *Node) error {
	node.NodeID = body.NodeId
	node.NodeIds = body.NodeIds

	initOk, err := NewMessage(
		msg.Dest,
		msg.Src,
		InitResponseBody{
			initOkType,
			body.MsgId,
		})

	if err != nil {
		return err
	}

	return node.sendMessage(initOk)
}

func handleEcho(msg Message, body RequestBody, node *Node) error {
	echoOk, err := NewMessage(
		msg.Dest,
		msg.Src,
		EchoResponseBody{
			echoOkType,
			body.MsgId,
			body.Echo,
		})

	if err != nil {
		return err
	}

	return node.sendMessage(echoOk)
}

func handleGenerate(msg Message, body RequestBody, node *Node) error {
	id := strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + msg.Dest
	generateOk, err := NewMessage(
		msg.Dest,
		msg.Src,
		GenerateResponseBody{
			generateOkType,
			body.MsgId,
			id,
		})
	if err != nil {
		return err
	}

	return node.sendMessage(generateOk)
}

func handleBroadcast(msg Message, body RequestBody, node *Node) error {
	if !slices.Contains(node.Messages, body.Message) {
		node.Messages = append(node.Messages, body.Message)
		for _, otherNode := range node.Topology[node.NodeID] {
			broadcast, err := NewMessage(node.NodeID, otherNode, body)
			if err != nil {
				return err
			}

			err = node.sendMessage(broadcast)
			if err != nil {
				return err
			}
		}
	}

	broadcastOk, err := NewMessage(
		msg.Dest,
		msg.Src,
		BroadcastResponseBody{
			broadcastOkType,
			body.MsgId,
		},
	)
	if err != nil {
		return err
	}

	return node.sendMessage(broadcastOk)
}

func handleRead(msg Message, body RequestBody, node *Node) error {
	readOk, err := NewMessage(
		msg.Dest,
		msg.Src,
		ReadResponseBody{
			readOkType,
			node.Messages,
			body.MsgId,
		},
	)
	if err != nil {
		return err
	}

	return node.sendMessage(readOk)
}

func handleTopology(msg Message, body RequestBody, node *Node) error {
	node.Topology = body.Topology

	topologyOk, err := NewMessage(
		msg.Dest,
		msg.Src,
		TopologyResponseBody{
			topologyOkType,
			body.MsgId,
		},
	)
	if err != nil {
		return err
	}

	return node.sendMessage(topologyOk)
}
