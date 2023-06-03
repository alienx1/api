package key

import (
	"encoding/hex"
)

func EncodeHex(str string) string {
	encoded := hex.EncodeToString([]byte(str))
	return encoded
}

func DecodeHex(hexStr string) string {
	decoded, _ := hex.DecodeString(hexStr)

	return string(decoded)
}
