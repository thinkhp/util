package thinkCrypto

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
	"util/think"
)

func TestEncoding(t *testing.T) {
	// utf-8
	str := "你好"
	b := []byte(str)
	fmt.Println(b, str)

	// hex
	str = hex.EncodeToString(b)
	fmt.Println(b, str)
}

func TestMd5(t *testing.T) {
	s := "15891467397"
	fmt.Println(GetMD5(s))
	fmt.Println(GetMD5Std(s))
}

var privateKeyStr = "MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAtd6mb1SR+VvacGr5sbEz3m5iWcqmNeJipJaGnJ5bGDErjglqLkVPXIRCbvMRUNEe/IlJdmLRT0sBzYJRDYQbcwIDAQABAkAPdFEOStBwsRZ50Q1QxS8UKqse2DKRh6A8PjJIIsi44GgYMvqXvlN+Vy5q5nYhkvB3Ndfhtn17f5qMalmRUlYRAiEA1/jXOGF4IWMF2okLEX6uRdN7J0o2iF8pcJGVr+l1gKkCIQDXk8SFzDUmZ7Ihvvns+NatHYx/U14Dnh5wdR04HgoguwIgTOAmu8r2F+xHiSJ+7htJrVE55SJlhuVYutkXjyZqzQECIQCYTWhxUqVWPbqGxuLRfbhFU/P33JE2IxbEQqljBS4IkwIgJswLDvDUHtjfuf0hLw61HIRN3ZflB7R/gj+Hw5qUspQ="
var publicKeyStr = "MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALXepm9Ukflb2nBq+bGxM95uYlnKpjXiYqSWhpyeWxgxK44Jai5FT1yEQm7zEVDRHvyJSXZi0U9LAc2CUQ2EG3MCAwEAAQ=="

func TestRSA(t *testing.T) {
	origData := "hello world"
	md5Data := md5.Sum([]byte(origData))
	fmt.Println("md5Data", base64.StdEncoding.EncodeToString(md5Data[:]))
	pubKey, err := GenPublicKey(publicKeyStr)
	think.IsNil(err)
	priKey, err := GenPrivateKeyPKCS8(privateKeyStr)
	think.IsNil(err)
	fmt.Println(priKey.PublicKey)
	fmt.Println(*pubKey)
	//
	signStd, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.MD5, md5Data[:])
	think.IsNil(err)

	signMy, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, md5Data[:])
	think.IsNil(err)
	parseData, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, signMy)
	think.IsNil(err)
	fmt.Println("parseData", base64.StdEncoding.EncodeToString(parseData))

	fmt.Println(base64.StdEncoding.EncodeToString(signStd))
	fmt.Println(base64.StdEncoding.EncodeToString(signMy))
	fmt.Println(string(signMy) == string(signStd))
}

//RSA加密、签名区别
//加密和签名都是为了安全性考虑，但略有不同。常有人问加密和签名是用私钥还是公钥？其实都是对加密和签名的作用有所混淆。简单的说，加密是为了防止信息被泄露，而签名是为了防止信息被篡改。这里举2个例子说明。
//
//第一个场景：战场上，B要给A传递一条消息，内容为某一指令。
//
//RSA的加密过程如下：
//
//（1）A生成一对密钥（公钥和私钥），私钥不公开，A自己保留。公钥为公开的，任何人可以获取。
//
//（2）A传递自己的公钥给B，B用A的公钥对消息进行加密。
//
//（3）A接收到B加密的消息，利用A自己的私钥对消息进行解密。
//
//　　在这个过程中，只有2次传递过程，第一次是A传递公钥给B，第二次是B传递加密消息给A，即使都被敌方截获，也没有危险性，因为只有A的私钥才能对消息进行解密，防止了消息内容的泄露。
//
//
//
//第二个场景：A收到B发的消息后，需要进行回复“收到”。
//
//RSA签名的过程如下：
//
//（1）A生成一对密钥（公钥和私钥），私钥不公开，A自己保留。公钥为公开的，任何人可以获取。
//
//（2）A用自己的私钥对消息加签，形成签名，并将加签的消息和消息本身一起传递给B。
//
//（3）B收到消息后，在获取A的公钥进行验签，如果验签出来的内容与消息本身一致，证明消息是A回复的。
//
//　　在这个过程中，只有2次传递过程，第一次是A传递加签的消息和消息本身给B，第二次是B获取A的公钥，即使都被敌方截获，也没有危险性，因为只有A的私钥才能对消息进行签名，即使知道了消息内容，也无法伪造带签名的回复给B，防止了消息内容的篡改。
//
//
//
//　　但是，综合两个场景你会发现，第一个场景虽然被截获的消息没有泄露，但是可以利用截获的公钥，将假指令进行加密，然后传递给A。第二个场景虽然截获的消息不能被篡改，但是消息的内容可以利用公钥验签来获得，并不能防止泄露。所以在实际应用中，要根据情况使用，也可以同时使用加密和签名，比如A和B都有一套自己的公钥和私钥，当A要给B发送消息时，先用B的公钥对消息加密，再对加密的消息使用A的私钥加签名，达到既不泄露也不被篡改，更能保证消息的安全性。
//
//　　总结：公钥加密、私钥解密、私钥签名、公钥验签。
