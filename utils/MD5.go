package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5En(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	sum := hash.Sum(nil) //'jiayan'
	s := hex.EncodeToString(sum)
	return s
}
