package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	initType     = "init"
	initOkType   = "init_ok"
	echoType     = "echo"
	echoOkType   = "echo_ok"
	nodeIdKey    = "node_id"
	typeKey      = "type"
	msgIdKey     = "msg_id"
	inReplyToKey = "in_reply_to"
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
			log.Printf("Error unmarshaling JSON: %s", err)
		}

		msgType, ok := received.Body[typeKey].(string)
		if !ok {
			return
		}

		switch msgType {
		case initType:
			nodeId, ok = received.Body[nodeIdKey].(string)
			if !ok {
				log.Printf("No nodeId in the message: %s", received)
				return
			}
			replyToInit(received)
		case echoType:
			replyToEcho(received)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %s", err)
	}
}

func replyToInit(msg Message) {
	responseBody := make(map[string]any)
	msgId, ok := msg.Body[msgIdKey].(float64)
	if !ok {
		log.Printf("No messageId in the message: %s", msg)
		return
	}
	responseBody[inReplyToKey] = msgId
	responseBody[typeKey] = initOkType

	var response Message
	response.Src = nodeId
	response.Dest = msg.Src
	response.Body = responseBody

	jsonResponse, _ := json.Marshal(response)
	fmt.Println(string(jsonResponse))
}

func replyToEcho(msg Message) {
	responseBody := make(map[string]any)
	msgId, ok := msg.Body[msgIdKey].(float64)
	if !ok {
		log.Printf("No messageId in the message: %s", msg)
		return
	}
	responseBody[inReplyToKey] = msgId
	responseBody[typeKey] = echoOkType
	echo, ok := msg.Body[echoType].(string)
	if !ok {
		log.Printf("No echo in the message: %s", msg)
		return
	}
	responseBody[echoType] = echo

	var response Message
	response.Src = nodeId
	response.Dest = msg.Src
	response.Body = responseBody

	jsonResponse, _ := json.Marshal(response)
	fmt.Println(string(jsonResponse))
}
