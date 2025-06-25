package main

import (
	"log"
	"os"
)

func main() {
	node := NewNode(
		os.Stdin,
		os.Stdout,
	)

	if err := node.run(); err != nil {
		log.Fatalf("Node error: %+v", err)
	}
}
