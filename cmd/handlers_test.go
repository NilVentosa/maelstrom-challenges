package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	client1 = "c1"
	node1   = "n1"
	msgId   = 1
	echo    = "hello there this should be echoed"
)

func TestGetReplyToInit(t *testing.T) {
	inputMsg, _ := newMessage(
		client1,
		node1,
		RequestBody{
			MsgId: msgId,
			Type:  initType,
		},
	)
	expectedReply, _ := newMessage(
		node1,
		client1,
		InitResponseBody{
			InReplyTo: msgId,
			Type:      initOkType,
		},
	)

	actualReply, _ := getReplyToInit(inputMsg)

	assertMessageEquals(t, expectedReply, actualReply)
}

func TestGetReplyToEcho(t *testing.T) {
	inputMsg, _ := newMessage(
		client1,
		node1,
		RequestBody{
			MsgId: msgId,
			Type:  echoType,
			Echo:  echo,
		},
	)
	expectedReply, _ := newMessage(
		node1,
		client1,
		EchoResponseBody{
			Type:      echoOkType,
			InReplyTo: msgId,
			Echo:      echo,
		},
	)

	actualReply, _ := getReplyToEcho(inputMsg)

	assertMessageEquals(t, expectedReply, actualReply)
}

func TestGetReplyToGenerate(t *testing.T) {
	inputMsg, _ := newMessage(
		client1,
		node1,
		RequestBody{
			MsgId: msgId,
			Type:  generateType,
		},
	)
	expectedReply, _ := newMessage(
		node1,
		client1,
		GenerateResponseBody{
			Type:      generateOkType,
			InReplyTo: msgId,
			Id:        "",
		},
	)

	actualReply, _ := getReplyToGenerate(inputMsg)
	var unmarshaledBody GenerateResponseBody
	json.Unmarshal(actualReply.Body, &unmarshaledBody)

	actualReply2, _ := getReplyToGenerate(inputMsg)
	var unmarshaledBody2 GenerateResponseBody
	json.Unmarshal(actualReply2.Body, &unmarshaledBody2)

	assertDestSrcEquals(t, expectedReply, actualReply)
	assert.NotEqual(t, unmarshaledBody.Id, unmarshaledBody2.Id)
}

func assertDestSrcEquals(t *testing.T, expected Message, actual Message) {
	assert.Equal(t, expected.Src, actual.Src)
	assert.Equal(t, expected.Dest, actual.Dest)
}

func assertMessageEquals(t *testing.T, expected Message, actual Message) {
	assertDestSrcEquals(t, expected, actual)
	assert.Equal(t, expected.Body, actual.Body)
}
