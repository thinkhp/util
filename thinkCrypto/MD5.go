package thinkCrypto

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	// cipherStr 格式[]byte
	cipherStr := md5Ctx.Sum(nil)
	// 16进制字符表示的编解码
	return hex.EncodeToString(cipherStr)
}
