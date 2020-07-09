package app

import (
	"crypto/md5"
	"encoding/hex"
)

func PasswordHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
