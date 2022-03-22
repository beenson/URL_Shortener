package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var RandFunc func(int) int

func Init() {
	rand.Seed(time.Now().Unix())

	RandFunc = func(max int) int {
		return rand.Intn(max)
	}
}

func GenerateCode(codeLen int) string {
	b := make([]byte, codeLen)
	for i := range b {
		b[i] = alphabet[RandFunc(len(alphabet))]
	}
	return string(b)
}
