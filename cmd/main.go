package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	initType   = "init"
	initOkType = "init_ok"
	echoType   = "echo"
	echoOkType = "echo_ok"

	echoKey      = "echo"
	nodeIdKey    = "node_id"
	typeKey      = "type"
	msgIdKey     = "msg_id"
	inReplyToKey = "in_reply_to"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var msg Message
		err := json.Unmarshal([]byte(scanner.Text()), &msg)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %+v", err)
		}

		msgType, ok := msg.Body[typeKey].(string)
		if !ok {
			return
		}

		var nodeId string
		switch msgType {
		case initType:
			nodeId, ok = msg.Body[nodeIdKey].(string)
			if !ok {
				log.Printf("No nodeId in the message: %+v", msg)
				return
			}
			fmt.Println(getReplyToInit(msg, nodeId))
		case echoType:
			fmt.Println(getReplyToEcho(msg, nodeId))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %+v", err)
	}
}
