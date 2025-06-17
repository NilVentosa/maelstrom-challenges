package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_Run(t *testing.T) {
	inputJSON := `
{"src":"c1","dest":"n1","body":{"type":"echo","msg_id":1,"echo":"Please echo this"}}
{"src":"c2","dest":"n1","body":{"type":"init","msg_id":2}}
`
	expectedOutput := `
{"src":"n1","dest":"c1","body":{"type":"echo_ok","in_reply_to":1,"echo":"Please echo this"}}
{"src":"n1","dest":"c2","body":{"type":"init_ok","in_reply_to":2}}
`
	in := strings.NewReader(strings.TrimSpace(inputJSON))
	out := new(bytes.Buffer)

	node := Node{in, out}

	err := node.run()
	if err != nil {
		t.Fatalf("node.Run() returned an unexpected error: %v", err)
	}

	got := strings.TrimSpace(out.String())
	expected := strings.TrimSpace(expectedOutput)

	assert.Equal(t, expected, got)
}

func TestHandleMessage_KnownType(t *testing.T) {
	input := []byte(`{"src":"c2","dest":"n1","body":{"type":"init","msg_id":2}}`)

	response, err := handleMessage(input)
	assert.Nil(t, err)

	var responseMessge Message
	var responseBody InitResponseBody
	json.Unmarshal(response, &responseMessge)
	err = json.Unmarshal(responseMessge.Body, &responseBody)

	assert.Nil(t, err)
	assert.Equal(t, initOkType, responseBody.Type)
}

func TestHandleMessage_UnknownType(t *testing.T) {
	input := []byte(`{"src":"c1","dest":"n1","body":{"type":"unknown_type"}}`)

	_, err := handleMessage(input)

	assert.True(t, strings.Contains(err.Error(), "unknown"))
}

func TestGetReply_BroadcastMessage(t *testing.T) {
	body := RequestBody{
		Type: broadcastType,
	}
	marshaledBody, _ := json.Marshal(body)
	message := Message{
		Dest: "destination",
		Src:  "source",
		Body: marshaledBody,
	}

	response, err := getReply(message)
	assert.Nil(t, err)

	var responseBody BroadcastResponseBody
	json.Unmarshal(response.Body, &responseBody)

	assert.Equal(t, responseBody.Type, broadcastOkType)
}

func TestGetReply_ReadMessage(t *testing.T) {
	body := RequestBody{
		Type: readType,
	}
	marshaledBody, _ := json.Marshal(body)
	message := Message{
		Dest: "destination",
		Src:  "source",
		Body: marshaledBody,
	}

	response, err := getReply(message)
	assert.Nil(t, err)

	var responseBody ReadResponseBody
	json.Unmarshal(response.Body, &responseBody)

	assert.Equal(t, responseBody.Type, readOkType)
}

func TestGetReply_TopologyMessage(t *testing.T) {
	body := RequestBody{
		Type: topologyType,
	}
	marshaledBody, _ := json.Marshal(body)
	message := Message{
		Dest: "destination",
		Src:  "source",
		Body: marshaledBody,
	}

	response, err := getReply(message)
	assert.Nil(t, err)

	var responseBody TopologyResponseBody
	json.Unmarshal(response.Body, &responseBody)

	assert.Equal(t, responseBody.Type, topologyOkType)
}
