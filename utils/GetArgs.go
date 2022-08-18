package utils

// GetArgs Returns the arguments of a command as a slice of bytes.
func GetArgs(args []string) [][]byte {
	var b [][]byte
	for _, v := range args {
		b = append(b, []byte(v))
	}
	return b
}
