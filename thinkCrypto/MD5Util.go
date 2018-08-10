package thinkCrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"kj/util"
	"util/think"
)

func GetMD5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 微信:
// 接口如果涉及敏感数据（如wx.getUserInfo当中的 openId 和unionId ），接口的明文内容将不包含这些敏感数据。
// 开发者如需要获取敏感数据，需要对接口返回的加密数据( encryptedData )进行对称解密。
// 解密算法如下：
// 对称解密使用的算法为 AES-128-CBC，数据采用PKCS#7填充。
// 对称解密的目标密文为 Base64_Decode(encryptedData)。
// 对称解密秘钥 aeskey = Base64_Decode(session_key), aeskey 是16字节。
// 对称解密算法初始向量 为Base64_Decode(iv)，其中iv由数据接口返回。
//
// 解密用户敏感数据获取用户信息
// sessionKey    数据进行加密签名的密钥
// encryptedData 包括敏感数据在内的完整用户信息的加密数据
// iv            加密算法的初始向量
func AESDecrypt(encryptedData, key, iv string) []byte {
	// 1.密文,密钥,初始向量 转为 []byte
	decode := func(str string) []byte {
		bytes, err := base64.StdEncoding.DecodeString(str)
		think.Check(err)
		return bytes
	}
	dataByte := decode(encryptedData)
	keyByte := decode(key)
	ivByte := decode(iv)

	// 2.保证密钥的len(byte)长度为16
	if len(keyByte)%aes.BlockSize != 0 {
		return nil
	}
	// 3.初始化
	block, err := aes.NewCipher(keyByte)
	util.IsNil(err)
	blockMode := cipher.NewCBCDecrypter(block, ivByte)
	// 4.解密
	origData := make([]byte, len(dataByte))
	blockMode.CryptBlocks(origData, dataByte)
	// 5.去PKCS#5填充
	// AES 中PKCS#5,PKCS#7 效果相同
	origData = PKCS5UnPadding(origData)

	return origData
}

func AESEncrypt(data, key, iv []byte) []byte {
	// 初始化
	block, err := aes.NewCipher(key)
	think.Check(err)
	mode := cipher.NewCBCEncrypter(block, iv)
	blockSize := block.BlockSize()
	// 数据采用PKCS#5填充
	data = PKCS5Padding(data, blockSize)
	// 加密
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(data, data)
	return data
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unPadding 次
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
