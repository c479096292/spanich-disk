package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}

// 密码加密
func EncodeMD5(val string) string {
	m :=md5.New()
	m.Write([]byte(val))
	return hex.EncodeToString(m.Sum(nil)) // 将加密后的byte转为string
}
