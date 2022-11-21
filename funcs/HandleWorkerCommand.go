package funcs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "kamigen/socket/structs"
	"reflect"
)

func HandleWorkerCommand(command *Command) {
	struc := GetStructure(command.Header.GetCommand())
	binary.Read(bytes.NewBuffer(command.Payload), binary.LittleEndian, struc)
	// Check if the structure has a proprety called "Player"
	typeOf := reflect.TypeOf(struc)
	_, b := typeOf.Elem().FieldByName("Player")
	if b {
		for client := range Clients {
			if Clients[client].Player == binary.BigEndian.Uint16(command.Payload) {
				Clients[client].Conn.Write(CreatePacket(command.Header.GetCommand(), WORKER, struc))
				return
			}
		}
		fmt.Println("Client not found")
	} else {
		fmt.Println("Structure does not implement Player")
	}
}

func GetStructure(packet Packet) interface{} {
	switch packet {
	case LOGIN:
		return &LoginRequest{}
	case LOGOUT:
		return &LogoutRequest{}
	case INFORMATION:
		return &InfoRequest{}
	case ROTATE:
		return &RotateRequest{}
	case CHAT:
		return &ChatRequest{}
	case SCREENSHOT:
		return &ScreenshotRequest{}
	}
	return nil
}
