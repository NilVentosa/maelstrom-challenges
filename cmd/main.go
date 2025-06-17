package main

import (
	"log"
	"os"
)

var nodeId string
var nodeIds string
var messages []any

func main() {
	node := Node{os.Stdin, os.Stdout}
	if err := node.run(); err != nil {
		log.Fatalf("Node error: %+v", err)
	}
}
