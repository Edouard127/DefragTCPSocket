package utils

import (
	"crypto/rand"
)

func GetRandomBytes(length int) []byte {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil
	}
	return b
}

func RandomUint16(length int) [16]byte {
	var b [16]byte
	random := GetRandomBytes(length)
	if random == nil {
		return b
	}
	copy(b[:], random)
	return b
}
