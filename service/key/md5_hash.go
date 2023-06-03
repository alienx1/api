package key

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMd5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}
