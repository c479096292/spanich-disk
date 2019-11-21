package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}

func MD5(b []byte) string {
	_md5 := md5.New()
	_md5.Write(b)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

// 密码加密
func EncodeMD5(val string) string {
	m :=md5.New()
	m.Write([]byte(val))
	return hex.EncodeToString(m.Sum(nil)) // 将加密后的byte转为string
}
