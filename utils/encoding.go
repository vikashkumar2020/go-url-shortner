package utils

import (
	"math"
	"strings"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func base62Encode(num int) string {
	encoded := ""
	for num > 0 {
		remainder := num % 62
		num /= 62
		encoded = string(base62Chars[remainder]) + encoded
	}
	return encoded
}
func base62Decode(str string) int {
	decoded := 0
	for i := 0; i < len(str); i++ {
		pos := strings.Index(base62Chars, string(str[i]))
		decoded += pos * int(math.Pow(62, float64(i)))
	}
	return decoded
}