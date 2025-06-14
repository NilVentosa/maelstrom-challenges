package main

import (
	"fmt"
)

func getReplyToInit(msg Message) (Message, error) {
	responseBody := make(map[string]any)
	msgId, ok := msg.Body[msgIdKey].(float64)
	if !ok {
		return Message{}, fmt.Errorf("No messageId in the message body: %+v", msg.Body)
	}
	responseBody[inReplyToKey] = msgId
	responseBody[typeKey] = initOkType

	return getReplyMessage(msg, responseBody)
}

func getReplyToEcho(msg Message) (Message, error) {
	responseBody := make(map[string]any)
	msgId, ok := msg.Body[msgIdKey].(float64)
	if !ok {
		return Message{}, fmt.Errorf("No messageId in the message body: %+v", msg.Body)
	}
	responseBody[inReplyToKey] = msgId
	responseBody[typeKey] = echoOkType
	echo, ok := msg.Body[echoType].(string)
	if !ok {
		return Message{}, fmt.Errorf("No echo in the message body: %+v", msg.Body)
	}
	responseBody[echoKey] = echo

	return getReplyMessage(msg, responseBody)
}

func getReplyMessage(msg Message, body map[string]any) (Message, error) {
	var response Message
	response.Src = msg.Dest
	response.Dest = msg.Src
	response.Body = body

	return response, nil
}
