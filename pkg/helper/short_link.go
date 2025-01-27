package helper

import (
	"strings"
)

func IdToShortURL(n uint64) string {
	charMap := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var shortURL strings.Builder

	for n > 0 {
		remainder := n % 62
		shortURL.WriteByte(charMap[remainder])
		n /= 62
	}

	runes := []rune(shortURL.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func ShortURLToID(shortURL string) uint64 {
	var id uint64

	for _, char := range shortURL {
		if char >= 'a' && char <= 'z' {
			id = id*62 + uint64(char-'a')
		} else if char >= 'A' && char <= 'Z' {
			id = id*62 + uint64(char-'A'+26)
		} else if char >= '0' && char <= '9' {
			id = id*62 + uint64(char-'0'+52)
		}
	}

	return id
}
