package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func handleMessage(msg Message) ([]byte, error) {
	response, responseError := getReplyToMessage(msg)

	if responseError != nil {
		return nil, responseError
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("problem marshaling the response: %w", err)
	}

	return jsonResponse, nil
}

func getReplyToMessage(msg Message) (Message, error) {
	var msgBody RequestBody
	if err := json.Unmarshal(msg.Body, &msgBody); err != nil {
		return Message{}, err
	}
	switch msgBody.Type {
	case initType:
		return getReplyToInit(msg)
	case echoType:
		return getReplyToEcho(msg)
	case generateType:
		return getReplyToGenerate(msg)
	case broadcastType:
		return getReplyToBroadcast(msg)
	case readType:
		return getReplyToRead(msg)
	case topologyType:
		return getReplyToTopology(msg)
	default:
		return Message{}, fmt.Errorf("unknown message type: %s", msgBody.Type)
	}
}
func getReplyToInit(msg Message) (Message, error) {
	var requestBody RequestBody
	json.Unmarshal(msg.Body, &requestBody)

	return NewMessage(
		msg.Dest,
		msg.Src,
		InitResponseBody{
			initOkType,
			requestBody.MsgId,
		})
}

func getReplyToEcho(msg Message) (Message, error) {
	var body RequestBody
	json.Unmarshal(msg.Body, &body)

	return NewMessage(
		msg.Dest,
		msg.Src,
		EchoResponseBody{
			echoOkType,
			body.MsgId,
			body.Echo,
		})
}

func getReplyToGenerate(msg Message) (Message, error) {
	var body RequestBody
	json.Unmarshal(msg.Body, &body)

	id := strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + msg.Dest
	return NewMessage(
		msg.Dest,
		msg.Src,
		GenerateResponseBody{
			generateOkType,
			body.MsgId,
			id,
		})
}

func getReplyToBroadcast(msg Message) (Message, error) {
	var body RequestBody
	json.Unmarshal(msg.Body, &body)

	messages = append(messages, body.Message)

	return NewMessage(
		msg.Dest,
		msg.Src,
		BroadcastResponseBody{
			broadcastOkType,
			body.MsgId,
		},
	)
}

func getReplyToRead(msg Message) (Message, error) {
	var body RequestBody
	json.Unmarshal(msg.Body, &body)

	return NewMessage(
		msg.Dest,
		msg.Src,
		ReadResponseBody{
			readOkType,
			messages,
			body.MsgId,
		},
	)
}

func getReplyToTopology(msg Message) (Message, error) {
	var body RequestBody
	json.Unmarshal(msg.Body, &body)

	return NewMessage(
		msg.Dest,
		msg.Src,
		TopologyResponseBody{
			topologyOkType,
			body.MsgId,
		},
	)
}
