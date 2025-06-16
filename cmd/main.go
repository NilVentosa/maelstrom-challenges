package main

import (
	"log"
	"os"
)

func main() {
	server := Server{os.Stdin, os.Stdout}
	if err := server.run(); err != nil {
		log.Fatalf("Server error: %+v", err)
	}
}
