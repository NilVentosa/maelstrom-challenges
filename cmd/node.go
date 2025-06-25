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
	Messages ConcurrentSet[int]
}

func NewNode(in io.Reader, out io.Writer) Node {
	return Node{
		In:       in,
		Out:      out,
		Messages: NewConcurrentSet[int](),
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
