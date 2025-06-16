package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func assertMessageEquals(t *testing.T, expected Message, actual Message) {
	assert.Equal(t, expected.Body, actual.Body)
	assert.Equal(t, expected.Dest, actual.Dest)
	assert.Equal(t, expected.Src, actual.Src)
}
