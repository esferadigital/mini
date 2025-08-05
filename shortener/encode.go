package shortener

import "strings"

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func EncodeBase62(num int64) string {
	if num == 0 {
		return "0"
	}

	var result strings.Builder
	for num > 0 {
		result.WriteByte(base62Chars[num%62])
		num /= 62
	}

	// Reverse the string since we built it backwards
	encoded := result.String()
	runes := []rune(encoded)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
