package main

const (
	messageTemplate  = `{"src":"%+v","dest":"%+v","body":%+v}`
	initBodyTemplate = `{
		"type":"init",
		"msg_id":%+v,
		"node_id":"%+v",
		"node_ids":["n1", "n2", "n3"]
	}`
	initOkBodyTemplate = `{
		"type":"init_ok",
		"in_reply_to":%+v
	}`
	errorBodyTemplate = `{
		"type":"error",
		"in_reply_to":5,
		"code":11,
		"text":"Node n5 is waiting for quorum and cannot service requests yet"
	}`
	echoBodyTemplate = `{
		"type":"echo",
		"msg_id":%+v,
		"echo":"%+v"
	}`
	echoOkBodyTemplate = `{
		"type":"echo_ok",
		"in_reply_to":%+v,
		"echo":"%+v"
	}`
	unknownTypeBody = `{
		"type":"unknown_type"
	}`
	generateBodyTemplate = `{
		"type":"generate",
		"msg_id":%+v
	}`
	generateOkBodyTemplate = `{
		"type": "generate_ok",
		"in_reply_to":%+v,
		"id":"%+v"
	}`
)
