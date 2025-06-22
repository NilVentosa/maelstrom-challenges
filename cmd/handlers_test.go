package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	client1 = "c1"
	node1   = "n1"
	msgId   = 1
	echo    = "hello there this should be echoed"
)

func TestHandleMessage_KnownType(t *testing.T) {
	msg, _ := newInitMessage(client1, node1, msgId, node1)

	node := newTestNode("")
	response, err := handleMessage(msg, &node)
	assert.Nil(t, err)

	var responseMessge Message
	var responseBody InitResponseBody
	json.Unmarshal(response, &responseMessge)
	err = json.Unmarshal(responseMessge.Body, &responseBody)

	assert.Nil(t, err)
	assert.Equal(t, initOkType, responseBody.Type)
}

func TestHandleMessage_UnknownType(t *testing.T) {
	msg, _ := newUnknownTypeMessage(client1, node1)
	node := newTestNode("")
	_, err := getReplyToMessage(msg, &node)

	assert.True(t, strings.Contains(err.Error(), "unknown"))
}

func TestGetReplyToMessage_Broadcast(t *testing.T) {
	body := RequestBody{
		Type: broadcastType,
	}
	marshaledBody, _ := json.Marshal(body)
	message := Message{
		Dest: node1,
		Src:  client1,
		Body: marshaledBody,
	}

	node := newTestNode("")
	response, err := getReplyToMessage(message, &node)
	assert.Nil(t, err)

	var responseBody BroadcastResponseBody
	json.Unmarshal(response.Body, &responseBody)

	assert.Equal(t, responseBody.Type, broadcastOkType)
}

func TestGetReplyToMessage_Init(t *testing.T) {
	initMessage, _ := newInitMessage(client1, node1, msgId, node1)
	expectedResponse, _ := newInitOkMessage(node1, client1, msgId)
	node := newTestNode("")

	actualResponse, _ := getReplyToMessage(initMessage, &node)

	assertMessageEquals(t, expectedResponse, actualResponse, &InitResponseBody{})
}

func TestGetReplyToMessage_Echo(t *testing.T) {
	echoMessage, _ := newEchoMessage(client1, node1, msgId, echo)
	expectedReply, _ := newEchoOkMessage(node1, client1, msgId, echo)
	node := newTestNode("")

	actualReply, _ := getReplyToMessage(echoMessage, &node)

	assertMessageEquals(t, expectedReply, actualReply, &EchoResponseBody{})
}

func TestGetReplyToMessage_Generate(t *testing.T) {
	inputMsg, _ := newGenerateMessage(client1, node1, msgId)
	expectedResponse, _ := newGenerateOkMessage(node1, client1, msgId, "")
	node := newTestNode("")

	actualReply, _ := getReplyToMessage(inputMsg, &node)
	var unmarshaledBody GenerateResponseBody
	json.Unmarshal(actualReply.Body, &unmarshaledBody)

	actualReply2, _ := getReplyToMessage(inputMsg, &node)
	var unmarshaledBody2 GenerateResponseBody
	json.Unmarshal(actualReply2.Body, &unmarshaledBody2)

	assertDestSrcEquals(t, expectedResponse, actualReply)
	assert.NotEqual(t, unmarshaledBody.Id, unmarshaledBody2.Id)
}

func TestGetReplyToMessage_Read(t *testing.T) {
	body := RequestBody{
		Type: readType,
	}
	marshaledBody, _ := json.Marshal(body)
	message := Message{
		Dest: node1,
		Src:  client1,
		Body: marshaledBody,
	}

	node := newTestNode("")
	response, err := getReplyToMessage(message, &node)
	assert.Nil(t, err)

	var responseBody ReadResponseBody
	json.Unmarshal(response.Body, &responseBody)

	assert.Equal(t, responseBody.Type, readOkType)
}

func TestGetReplyToMessage_Topology(t *testing.T) {
	body := RequestBody{
		Type: topologyType,
	}
	marshaledBody, _ := json.Marshal(body)
	message := Message{
		Dest: node1,
		Src:  client1,
		Body: marshaledBody,
	}

	node := newTestNode("")
	response, err := getReplyToMessage(message, &node)
	assert.Nil(t, err)

	var responseBody TopologyResponseBody
	json.Unmarshal(response.Body, &responseBody)

	assert.Equal(t, responseBody.Type, topologyOkType)
}

func assertDestSrcEquals(t *testing.T, expected Message, actual Message) {
	assert.Equal(t, expected.Src, actual.Src)
	assert.Equal(t, expected.Dest, actual.Dest)
}

func assertMessageEquals(t *testing.T, expected Message, actual Message, typ any) {
	assertDestSrcEquals(t, expected, actual)
	assertBodyEquals(t, expected.Body, actual.Body, typ)
}

func assertBodyEquals(t *testing.T, expected, actual json.RawMessage, typ any) {
	targetType := reflect.TypeOf(typ)

	expectedVal := reflect.New(targetType.Elem()).Interface()
	actualVal := reflect.New(targetType.Elem()).Interface()

	json.Unmarshal(expected, expectedVal)
	json.Unmarshal(actual, actualVal)

	assert.Equal(t, expectedVal, actualVal)
}
