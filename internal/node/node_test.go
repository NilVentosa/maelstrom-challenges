package node

import (
	"bytes"
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

	node := Node{
		In:  in,
		Out: out,
	}

	err := node.Run()
	if err != nil {
		t.Fatalf("node.Run() returned an unexpected error: %v", err)
	}

	got := strings.TrimSpace(out.String())
	expected := strings.TrimSpace(expectedOutput)

	assert.Equal(t, expected, got)
}
