package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Node struct {
	In       io.Reader
	Out      io.Writer
	NodeID   string
	NodeIds  []string
	Topology map[string][]string
	Messages ConcurrentSet[int]
	Pending  ConcurrentSet[PendingAck]
}

func NewNode(in io.Reader, out io.Writer) Node {
	return Node{
		In:       in,
		Out:      out,
		Messages: NewConcurrentSet[int](),
		Pending:  NewConcurrentSet[PendingAck](),
	}
}

func (node *Node) run() error {
	scanner := bufio.NewScanner(node.In)
	for scanner.Scan() {
		var msg Message
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			return fmt.Errorf("error unmarshaling message: %w", err)
		}

		err := handleMessage(msg, node)
		if err != nil {
			return fmt.Errorf("failed to handle message: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading from input: %w", err)
	}
	return nil
}

func (node *Node) sendMessage(message Message) error {
	jsonResponse, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(node.Out, string(jsonResponse))
	return err
}

func (node *Node) sendMessageUntilAck(message Message, body RequestBody) {
	// TODO: make the type of PendingAck more generic
	pendingAck := PendingAck{body.MsgID, broadcastOkType, message.Dest}
	node.Pending.Add(pendingAck)
	for node.Pending.Contains(pendingAck) {
		if err := node.sendMessage(message); err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
}
