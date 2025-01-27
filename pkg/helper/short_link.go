package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateShortLink(url string) string {
	hash := md5.Sum([]byte(url))
	hexString := hex.EncodeToString(hash[:])
	return hexString[:8]
}
