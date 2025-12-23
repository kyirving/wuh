package util

import (
	"crypto/md5"
	"fmt"
)

func EncryptPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
