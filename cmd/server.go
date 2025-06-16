package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

type Server struct {
	In  io.Reader
	Out io.Writer
}

func (server *Server) run() error {
	scanner := bufio.NewScanner(server.In)
	for scanner.Scan() {
		response, err := handleMessage(scanner.Bytes())
		if err != nil {
			return fmt.Errorf("failed to handle message: %w", err)
		}

		if _, err := fmt.Fprintln(server.Out, string(response)); err != nil {
			return fmt.Errorf("failed to write response: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading from input: %w", err)
	}
	return nil
}

func handleMessage(input []byte) ([]byte, error) {
	var msg Message
	if err := json.Unmarshal(input, &msg); err != nil {
		return nil, fmt.Errorf("error unmarshaling message: %w", err)
	}

	var msgBody RequestBody
	if err := json.Unmarshal(msg.Body, &msgBody); err != nil {
		return nil, fmt.Errorf("error unmarshaling message body: %w", err)
	}

	var response Message
	var responseError error

	switch msgBody.Type {
	case initType:
		response, responseError = getReplyToInit(msg)
	case echoType:
		response, responseError = getReplyToEcho(msg)
	case generateType:
		response, responseError = getReplyToGenerate(msg)
	default:
		responseError = fmt.Errorf("unknown message type: %s", msgBody.Type)
	}

	if responseError != nil {
		return nil, responseError
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("problem marshaling the response: %w", err)
	}

	return jsonResponse, nil
}
