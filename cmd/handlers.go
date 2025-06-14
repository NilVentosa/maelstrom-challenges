package main

import (
	"encoding/json"
	"strconv"
	"time"
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

func getReplyToGenerate(msg Message, body RequestBody) (Message, error) {
	epochTimestamp := time.Now().UnixNano()
	id := strconv.FormatInt(epochTimestamp, 10) + "-" + msg.Dest
	responseBody, err := json.Marshal(GenerateResponseBody{generateOkType, body.MsgId, id})
	if err != nil {
		return Message{}, err
	}
	responseMessage := Message{msg.Dest, msg.Src, responseBody}
	return responseMessage, nil
}
