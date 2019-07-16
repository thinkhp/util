package thinkCrypto

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"util/think"
	"util/thinkJson"
)

// 获取公钥
func GenPublicKey(publicKeyStr string) (pubKey *rsa.PublicKey, err error) {
	// BASE64解码(pemBlock)
	bs, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return nil, err
	}
	parseKey, err := x509.ParsePKIXPublicKey(bs) // not a public pem
	if err != nil {
		return nil, err
	}
	pubKey, ok := parseKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key is not a valid RSA public key")
	}
	//// cert
	//cert, err := x509.ParseCertificate(bs)
	//if err != nil {
	//	return nil, err
	//}
	//// public key
	//pubKey, ok := cert.PublicKey.(*rsa.PublicKey)
	//if !ok {
	//	return nil, fmt.Errorf("key is not a valid RSA public key")
	//}

	return pubKey, nil
}

// 获取私钥
func GenPrivateKeyPKCS8(privateKey string) (priKey *rsa.PrivateKey, err error) {
	bs, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(bs)
	if err != nil {
		return nil, err
	}

	priKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("key is not a valid RSA private key")
	}

	return priKey, err
}

// 签名
// md5WithRSA签名
func MakeSignMy(params map[string]interface{}, privateKeyStr string) string {
	// 以 PKCS8 格式生成私钥
	privateKey, err := GenPrivateKeyPKCS8(privateKeyStr)
	think.IsNil(err)
	// data 字典排序输出 json
	jsonSorted := thinkJson.MapJsonBySortKey(params, nil)
	// 内容摘要 md5
	sum := md5.Sum([]byte(jsonSorted))
	signByte, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, sum[:])
	think.IsNil(err)
	// base64
	return base64.StdEncoding.EncodeToString(signByte)
}

func VerifySignMy(params map[string]interface{}, publicKeyStr, sign string) bool{
	//
	publicKey, err := GenPublicKey(publicKeyStr)
	think.IsNil(err)

	// data 字典排序输出 json
	jsonSorted := thinkJson.MapJsonBySortKey(params, nil)
	// 内容摘要 md5
	sum := md5.Sum([]byte(jsonSorted))
	sumEn, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(sign))
	think.IsNil(err)
	return bytes.Equal(sum[:], sumEn)
}


func UnPadding(src []byte, keySize int) [][]byte {

	srcSize := len(src)

	blockSize := keySize

	var v [][]byte

	if srcSize == blockSize {
		v = append(v, src)
	} else {
		groups := len(src) / blockSize
		for i := 0; i < groups; i++ {
			block := src[:blockSize]

			v = append(v, block)
			src = src[blockSize:]
		}
	}
	return v
}

