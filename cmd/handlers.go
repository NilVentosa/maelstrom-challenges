package main

import (
	"encoding/json"
	"strconv"
	"time"
)

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
