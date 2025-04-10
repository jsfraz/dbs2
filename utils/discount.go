package utils

import (
	"math/rand"
)

// Funkce generující náhodný alfanumerický řetězec.
//
//	@return string
func GenerateRandomString(length int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
