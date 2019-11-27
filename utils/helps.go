package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"os/exec"
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


// 执行 linux shell command
func ExecLinuxShell(s string) (string, error) {
	//函数返回一个io.Writer类型的*Cmd
	cmd := exec.Command("/bin/bash", "-c", s)

	//通过bytes.Buffer将byte类型转化为string类型
	var result bytes.Buffer
	cmd.Stdout = &result

	//Run执行cmd包含的命令，并阻塞直至完成
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return result.String(), err
}