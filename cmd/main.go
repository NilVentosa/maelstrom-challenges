package main

import (
	"log"
	"os"
)

var nodeId string
var nodeIds string
var messages []any

func main() {
	node := Node{
		In:  os.Stdin,
		Out: os.Stdout,
	}
	if err := node.run(); err != nil {
		log.Fatalf("Node error: %+v", err)
	}
}
