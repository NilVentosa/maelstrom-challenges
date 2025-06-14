package main

import (
	"encoding/json"
	"fmt"
)

func getReplyToInit(msg Message, nodeId string) (string, error) {
	responseBody := make(map[string]any)
	msgId, ok := msg.Body[msgIdKey].(float64)
	if !ok {
		return "", fmt.Errorf("No messageId in the message body: %+v", msg.Body)
	}
	responseBody[inReplyToKey] = msgId
	responseBody[typeKey] = initOkType

	return getReplyMessage(msg, responseBody, nodeId)
}

func getReplyToEcho(msg Message, nodeId string) (string, error) {
	responseBody := make(map[string]any)
	msgId, ok := msg.Body[msgIdKey].(float64)
	if !ok {
		return "", fmt.Errorf("No messageId in the message body: %+v", msg.Body)
	}
	responseBody[inReplyToKey] = msgId
	responseBody[typeKey] = echoOkType
	echo, ok := msg.Body[echoType].(string)
	if !ok {
		return "", fmt.Errorf("No echo in the message body: %+v", msg.Body)
	}
	responseBody[echoKey] = echo

	return getReplyMessage(msg, responseBody, nodeId)
}

func getReplyMessage(msg Message, body map[string]any, nodeId string) (string, error) {
	var response Message
	response.Src = nodeId
	response.Dest = msg.Src
	response.Body = body

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(jsonResponse), nil
}
