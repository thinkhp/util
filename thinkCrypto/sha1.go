package thinkCrypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func GetSha1(str string) string {
	sum := sha1.Sum([]byte(str))
	// 与上一句的同等实现,摘自标准库
	//s := sha1.New()
	//s.Reset()
	//s.Write([]byte(str))
	//sum := s.Sum(nil)

	return hex.EncodeToString(sum[:])
}
