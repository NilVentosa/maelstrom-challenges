package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func replyToInit(msg Message) {
	responseBody := make(map[string]any)
	msgId, ok := msg.Body[msgIdKey].(float64)
	if !ok {
		log.Printf("No messageId in the message: %+v", msg)
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
		log.Printf("No messageId in the message: %+v", msg)
		return
	}
	responseBody[inReplyToKey] = msgId
	responseBody[typeKey] = echoOkType
	echo, ok := msg.Body[echoType].(string)
	if !ok {
		log.Printf("No echo in the message: %+v", msg)
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
