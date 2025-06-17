package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

type Node struct {
	In       io.Reader
	Out      io.Writer
	NodeID   string
	NodeIds  []string
	Topology map[string][]string
}

func (node *Node) run() error {
	scanner := bufio.NewScanner(node.In)
	for scanner.Scan() {
		var msg Message
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			return fmt.Errorf("error unmarshaling message: %w", err)
		}
		var msgBody RequestBody
		if err := json.Unmarshal(msg.Body, &msgBody); err != nil {
			return err
		}

		if msgBody.Type == initType {
			node.NodeID = msgBody.NodeId
			node.NodeIds = msgBody.NodeIds
		}
		if msgBody.Type == topologyType {
			node.Topology = msgBody.Topology
		}
		response, err := handleMessage(msg)
		if err != nil {
			return fmt.Errorf("failed to handle message: %w", err)
		}

		if _, err := fmt.Fprintln(node.Out, string(response)); err != nil {
			return fmt.Errorf("failed to write response: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading from input: %w", err)
	}
	return nil
}
