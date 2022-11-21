package structs

import "encoding/binary"

type Fragment struct {
	Checksum [16]byte
	Payload  []byte
}

func (f *Fragment) GetChecksum() uint16 {
	return binary.LittleEndian.Uint16(f.Checksum[:])
}
