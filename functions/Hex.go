package functions

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hex(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
