package main

import (
	"log"
	"os"

	"github.com/nilventosa/maelstrom-challenges/internal/node"
)

func main() {
	node := node.NewNode(
		os.Stdin,
		os.Stdout,
	)

	if err := node.Run(); err != nil {
		log.Fatalf("Node error: %+v", err)
	}
}
