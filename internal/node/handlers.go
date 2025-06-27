package node

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/nilventosa/maelstrom-challenges/internal/messages"
)

func handleMessage(msg messages.Message, node *Node) error {
	var body messages.RequestBody
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	switch body.Type {
	case messages.InitType:
		return handleInit(msg, body, node)
	case messages.EchoType:
		return handleEcho(msg, body, node)
	case messages.GenerateType:
		return handleGenerate(msg, body, node)
	case messages.BroadcastType:
		return handleBroadcast(msg, body, node)
	case messages.BroadcastOkType:
		return handleBroadcastOk(msg, body, node)
	case messages.ReadType:
		return handleRead(msg, body, node)
	case messages.TopologyType:
		return handleTopology(msg, body, node)
	default:
		return fmt.Errorf("unknown message type: %s", body.Type)
	}
}

func handleInit(msg messages.Message, body messages.RequestBody, node *Node) error {
	node.NodeID = body.NodeID
	node.NodeIds = body.NodeIds

	initOk, err := messages.NewMessage(
		msg.Dest,
		msg.Src,
		messages.InitResponseBody{
			Type:      messages.InitOkType,
			InReplyTo: body.MsgID,
		})
	if err != nil {
		return err
	}

	return node.sendMessage(initOk)
}

func handleEcho(msg messages.Message, body messages.RequestBody, node *Node) error {
	echoOk, err := messages.NewMessage(
		msg.Dest,
		msg.Src,
		messages.EchoResponseBody{
			Type:      messages.EchoOkType,
			InReplyTo: body.MsgID,
			Echo:      body.Echo,
		})
	if err != nil {
		return err
	}

	return node.sendMessage(echoOk)
}

func handleGenerate(msg messages.Message, body messages.RequestBody, node *Node) error {
	id := strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + msg.Dest
	generateOk, err := messages.NewMessage(
		msg.Dest,
		msg.Src,
		messages.GenerateResponseBody{
			Type:      messages.GenerateOkType,
			InReplyTo: body.MsgID,
			ID:        id,
		})
	if err != nil {
		return err
	}

	return node.sendMessage(generateOk)
}

func handleBroadcast(msg messages.Message, body messages.RequestBody, node *Node) error {
	if !node.Messages.Contains(body.Message) {
		node.Messages.Add(body.Message)
		for _, n := range node.Topology[node.NodeID] {
			broadcast, err := messages.NewMessage(node.NodeID, n, body)
			if err != nil {
				return err
			}

			go node.sendMessageUntilAck(broadcast, body)
		}
	}

	broadcastOk, err := messages.NewMessage(
		msg.Dest,
		msg.Src,
		messages.BroadcastResponseBody{
			Type:      messages.BroadcastOkType,
			InReplyTo: body.MsgID,
		},
	)
	if err != nil {
		return err
	}

	return node.sendMessage(broadcastOk)
}

func handleBroadcastOk(msg messages.Message, body messages.RequestBody, node *Node) error {
	node.Pending.Remove(
		PendingAck{
			body.InReplyTo,
			messages.BroadcastOkType,
			msg.Src,
		})
	return nil
}

func handleRead(msg messages.Message, body messages.RequestBody, node *Node) error {
	readOk, err := messages.NewMessage(
		msg.Dest,
		msg.Src,
		messages.ReadResponseBody{
			Type:      messages.ReadOkType,
			Messages:  node.Messages.Values(),
			InReplyTo: body.MsgID,
		},
	)
	if err != nil {
		return err
	}

	return node.sendMessage(readOk)
}

func handleTopology(msg messages.Message, body messages.RequestBody, node *Node) error {
	node.Topology = body.Topology

	topologyOk, err := messages.NewMessage(
		msg.Dest,
		msg.Src,
		messages.TopologyResponseBody{
			Type:      messages.TopologyOkType,
			InReplyTo: body.MsgID,
		},
	)
	if err != nil {
		return err
	}

	return node.sendMessage(topologyOk)
}
