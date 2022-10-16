package cf

import (
	"crypto/md5"
	"encoding/hex"
)

// 计算MD5
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}
