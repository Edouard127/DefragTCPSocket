package utils

// ByteArraysExtract Makes an array of bytes from an array of array of bytes.
func ByteArraysExtract(arr [][]byte) []byte {
	var b []byte
	for _, v := range arr {
		b = append(b, v...)
		b = append(b, ' ')
	}
	return b
}
