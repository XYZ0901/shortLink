package utils

import (
	"crypto/md5"
	"fmt"
)

func ToHashString(s string) (string, error) {
	hash := md5.New()
	if _, err := hash.Write([]byte(s)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", hash.Sum(nil)), nil
}
