package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Message struct {
	Src  string         `json:"src"`
	Dest string         `json:"dest"`
	Body map[string]any `json:"body"`
}

var nodeId string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var received Message
		err := json.Unmarshal([]byte(scanner.Text()), &received)
		if err != nil {
			log.Fatalf("Error unmarshaling JSON: %s", err)
		}

		switch received.Body["type"].(string) {
		case "init":
			nodeId = received.Body["node_id"].(string)
			replyToInit(received)
		case "echo":
			replyToEcho(received)
		}
	}

	if err := scanner.Err(); err != nil {
	}
}

func replyToInit(received Message) {
	responseBody := make(map[string]any)
	responseBody["in_reply_to"] = received.Body["msg_id"].(float64)
	responseBody["type"] = "init_ok"

	var response Message
	response.Src = nodeId
	response.Dest = received.Src
	response.Body = responseBody

	json_response, _ := json.Marshal(response)
	fmt.Println(string(json_response))
}

func replyToEcho(received Message) {
	responseBody := make(map[string]any)
	responseBody["in_reply_to"] = received.Body["msg_id"].(float64)
	responseBody["type"] = "echo_ok"
	responseBody["echo"] = received.Body["echo"].(string)

	var response Message
	response.Src = nodeId
	response.Dest = received.Src
	response.Body = responseBody

	jsonResponse, _ := json.Marshal(response)
	fmt.Println(string(jsonResponse))
}
