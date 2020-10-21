package utils

import (
	"crypto/md5"
	"fmt"
)

// MD5Hash is a function to hash string data
func MD5Hash(src string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(src)))
}
