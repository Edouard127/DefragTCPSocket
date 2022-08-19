package utils

import "crypto/rand"

func GetRandomBytes(length int) []byte {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil
	}
	return b
}
