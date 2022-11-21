package structs

import (
	"bytes"
	"encoding/binary"
)

type Packet uint8

const (
	EXIT Packet = iota
	OK
	HEARTBEAT
	LOGIN
	LOGOUT
	ADD_WORKER
	REMOVE_WORKER
	ADD_LISTENER
	REMOVE_LISTENER
	INFORMATION
	ROTATE
	HIGHWAY_TOOLS
	JOB
	GET_JOBS
	CHAT
	BARITONE
	LAMBDA
	ERROR
	SCREENSHOT
)

func (p Packet) String() string {
	// Return the string representation of the constant
	switch p {
	case EXIT:
		return "EXIT"
	case OK:
		return "OK"
	case HEARTBEAT:
		return "HEARTBEAT"
	case LOGIN:
		return "LOGIN"
	case LOGOUT:
		return "LOGOUT"
	case ADD_WORKER:
		return "ADD_WORKER"
	case REMOVE_WORKER:
		return "REMOVE_WORKER"
	case ADD_LISTENER:
		return "ADD_LISTENER"
	case REMOVE_LISTENER:
		return "REMOVE_LISTENER"
	case INFORMATION:
		return "INFORMATION"
	case ROTATE:
		return "ROTATE"
	case HIGHWAY_TOOLS:
		return "HIGHWAY_TOOLS"
	case JOB:
		return "JOB"
	case GET_JOBS:
		return "GET_JOBS"
	case CHAT:
		return "CHAT"
	case BARITONE:
		return "BARITONE"
	case LAMBDA:
		return "LAMBDA"
	case ERROR:
		return "ERROR"
	case SCREENSHOT:
		return "SCREENSHOT"
	}
	return "UNKNOWN"
}

func CreatePacket(command Packet, destination Destination, payload interface{}) []byte {
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.LittleEndian, payload); err != nil {
		return nil
	}
	header := CommandHeader{
		Command:     [1]byte{byte(command)},
		Destination: [1]byte{byte(destination)},
	}
	c := Command{
		Header:  header,
		Payload: buffer.Bytes(),
	}
	nbuffer := new(bytes.Buffer)
	nbuffer.Write(c.Header.Command[:])
	nbuffer.Write(c.Header.Destination[:])
	nbuffer.Write(c.Payload)
	nbuffer.Write([]byte("\r\n"))
	return nbuffer.Bytes()
}
