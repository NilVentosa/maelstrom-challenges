package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	// Message types
	initType       = "init"
	initOkType     = "init_ok"
	echoType       = "echo"
	echoOkType     = "echo_ok"
	generateType   = "generate"
	generateOkType = "generate_ok"

	// Keys in the messages
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

		var msgBody RequestBody
		err = json.Unmarshal(msg.Body, &msgBody)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %+v", err)
		}

		var response Message
		var responseError error
		switch msgBody.Type {
		case initType:
			response, responseError = getReplyToInit(msg)
		case echoType:
			response, responseError = getReplyToEcho(msg)
		case generateType:
			response, responseError = getReplyToGenerate(msg)
		}
		if responseError != nil {
			log.Fatalf("There was an error: %+v", responseError)
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Problem marshalling the response: %+v", err)
		}
		fmt.Println(string(jsonResponse), nil)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %+v", err)
	}
}
