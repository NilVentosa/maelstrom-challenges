package main

import (
	"encoding/json"
)

func getReplyToInit(msg Message, body RequestBody) (Message, error) {
	responseBody, err := json.Marshal(InitResponseBody{initOkType, body.MsgId})
	if err != nil {
		return Message{}, err
	}
	responseMessage := Message{msg.Dest, msg.Src, responseBody}
	return responseMessage, nil
}

func getReplyToEcho(msg Message, body RequestBody) (Message, error) {
	responseBody, err := json.Marshal(EchoResponseBody{echoOkType, body.MsgId, body.Echo})
	if err != nil {
		return Message{}, err
	}
	responseMessage := Message{msg.Dest, msg.Src, responseBody}
	return responseMessage, nil
}
