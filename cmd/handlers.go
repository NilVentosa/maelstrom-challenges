package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type PendingAck struct {
	MsgID int
	Type  string
	From  string
}

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
		return handleBroadcastOk(msg, body, node)
	case readType:
		return handleRead(msg, body, node)
	case topologyType:
		return handleTopology(msg, body, node)
	default:
		return fmt.Errorf("unknown message type: %s", body.Type)
	}
}

func handleInit(msg Message, body RequestBody, node *Node) error {
	node.NodeID = body.NodeID
	node.NodeIds = body.NodeIds

	initOk, err := NewMessage(
		msg.Dest,
		msg.Src,
		InitResponseBody{
			initOkType,
			body.MsgID,
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
			body.MsgID,
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
			body.MsgID,
			id,
		})
	if err != nil {
		return err
	}

	return node.sendMessage(generateOk)
}

func handleBroadcast(msg Message, body RequestBody, node *Node) error {
	if !node.Messages.Contains(body.Message) {
		node.Messages.Add(body.Message)
		for _, n := range node.Topology[node.NodeID] {
			broadcast, err := NewMessage(node.NodeID, n, body)
			if err != nil {
				return err
			}

			go node.sendMessageUntilAck(broadcast, body)
		}
	}

	broadcastOk, err := NewMessage(
		msg.Dest,
		msg.Src,
		BroadcastResponseBody{
			broadcastOkType,
			body.MsgID,
		},
	)
	if err != nil {
		return err
	}

	return node.sendMessage(broadcastOk)
}

func handleBroadcastOk(msg Message, body RequestBody, node *Node) error {
	node.Pending.Remove(PendingAck{body.InReplyTo, broadcastOkType, msg.Src})
	return nil
}

func handleRead(msg Message, body RequestBody, node *Node) error {
	var messages []int
	for key := range node.Messages.Values() {
		messages = append(messages, key)
	}
	readOk, err := NewMessage(
		msg.Dest,
		msg.Src,
		ReadResponseBody{
			readOkType,
			messages,
			body.MsgID,
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
			body.MsgID,
		},
	)
	if err != nil {
		return err
	}

	return node.sendMessage(topologyOk)
}
