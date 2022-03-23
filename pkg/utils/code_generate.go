package util

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var RandFunc func(int) int

func GenerateCode(codeLen int) string {
	b := make([]byte, codeLen)
	for i := range b {
		b[i] = alphabet[RandFunc(len(alphabet))]
	}
	return string(b)
}
