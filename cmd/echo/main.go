package main

import (
	"bufio"
	"encoding/json"
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

var nodeId string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var received Message
		err := json.Unmarshal([]byte(scanner.Text()), &received)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %+v", err)
		}

		msgType, ok := received.Body[typeKey].(string)
		if !ok {
			return
		}

		switch msgType {
		case initType:
			nodeId, ok = received.Body[nodeIdKey].(string)
			if !ok {
				log.Printf("No nodeId in the message: %+v", received)
				return
			}
			replyToInit(received)
		case echoType:
			replyToEcho(received)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %+v", err)
	}
}
